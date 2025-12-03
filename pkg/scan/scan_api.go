// Copyright 2022 Praetorian Security, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scan

import (
	"fmt"
	"log"
	"net/netip"
	"sync"

	"github.com/AbelChe/fingerprintx_pro/pkg/plugins"
)

func UDPScan(targets []plugins.Target, config Config) ([]plugins.Service, error) {
	var results []plugins.Service
	for _, target := range targets {
		result, err := config.UDPScanTarget(target)
		if err == nil && result != nil {
			results = append(results, *result)
		}
		if config.Verbose && err != nil {
			log.Printf("%s\n", err)
		}
	}

	return results, nil
}

// ScanTargets fingerprints service(s) running given a list of targets.
func ScanTargets(targets []plugins.Target, config Config) ([]plugins.Service, error) {
	var results []plugins.Service

	if config.UDP {
		return UDPScan(targets, config)
	}

	for _, target := range targets {
		result, err := config.SimpleScanTarget(target)
		if err == nil && result != nil {
			results = append(results, *result)
		}
		if config.Verbose && err != nil {
			log.Printf("%s\n", err)
		}
	}

	return results, nil
}

func CreateTargets(ipList []string, portList []int) (chan plugins.Target, chan error) {
	targetCh := make(chan plugins.Target)
	errCh := make(chan error)

	go func() {
		for _, port := range portList {
			for _, ipStr := range ipList {
				ip, err := netip.ParseAddr(ipStr)
				if err != nil {
					errCh <- fmt.Errorf("error parsing IP %s: %v", ipStr, err)
					continue
				}
				targetCh <- plugins.Target{
					Address: netip.AddrPortFrom(ip, uint16(port)),
					Host:    ipStr,
				}
			}
		}
		close(targetCh)
		close(errCh)
	}()

	return targetCh, errCh
}

func ScanTargetsMax(targets []plugins.Target, config Config, maxConcurrent int) ([]plugins.Service, error) {
	servicesCh := make(chan plugins.Service, len(targets))
	errCh := make(chan error, len(targets))

	sem := make(chan struct{}, maxConcurrent) //创建信号量

	var wg sync.WaitGroup

	if config.UDP {
		wg.Add(1)
		go func() {
			defer wg.Done()
			services, err := UDPScan(targets, config)
			if err != nil {
				errCh <- err
			} else {
				for _, service := range services {
					servicesCh <- service
				}
			}
		}()
	} else {
		for _, target := range targets {
			wg.Add(1)
			go func(target plugins.Target) {
				sem <- struct{}{} // 试图向sem发送数据，如果sem已满，则此处会阻塞，从而限制并发
				defer func() {
					<-sem // 保证函数结束时释放sem，允许其他协程开始执行
					wg.Done()
				}()

				service, err := config.SimpleScanTarget(target)
				if err != nil {
					errCh <- err
				} else if service != nil {
					servicesCh <- *service
				}
			}(target)
		}
	}

	//确保所有goroutinue完成后关闭channels
	go func() {
		wg.Wait()
		close(servicesCh)
		close(errCh)
	}()

	var results []plugins.Service
	var err error
	for service := range servicesCh {
		results = append(results, service)
	}
	for err = range errCh {
		if err != nil && config.Verbose {
			log.Printf("%s\n", err)
		}
	}

	return results, err
}
