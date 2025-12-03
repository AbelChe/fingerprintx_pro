package main

import (
	"fmt"
	"time"

	"github.com/AbelChe/fingerprintx_pro/pkg/plugins"
	"github.com/AbelChe/fingerprintx_pro/pkg/scan"
)

func main() {

	fxConfig := scan.Config{
		DefaultTimeout: time.Duration(2) * time.Second,
		FastMode:       false,
		Verbose:        false,
		UDP:            false,
	}

	portList := []int{
		80,
		443,
		3306,
		3389,
		8443,
	}

	ipList := []string{
		"119.147.111.212",
		"119.147.111.193",
		"119.147.111.242",
		"119.147.111.135",
		"119.147.111.223",
		"119.147.111.229",
		"119.147.111.219",
		"119.147.111.227",
		"119.147.111.186",
		"119.147.111.214",
		"119.147.111.202",
		"119.147.111.209",
		"119.147.111.145",
		"119.147.111.189",
		"119.147.111.238",
		"119.147.111.234",
		"119.147.111.110",
		"119.147.111.188",
		"119.147.111.222",
		"119.147.111.249",
		"119.147.111.141",
		"119.147.111.187",
		"119.147.111.240",
		"119.147.111.143",
		"119.147.111.179",
		"119.147.111.241",
		"119.147.111.127",
		"119.147.111.217",
		"119.147.111.180",
		"119.147.111.192",
		"119.147.111.207",
		"119.147.111.122",
		"119.147.111.215",
		"119.147.111.230",
		"119.147.111.191",
		"119.147.111.111",
		"119.147.111.232",
		"119.147.111.231",
		"119.147.111.208",
		"119.147.111.148",
		"119.147.111.216",
		"119.147.111.201",
		"119.147.111.108",
		"119.147.111.144",
		"119.147.111.178",
		"119.147.111.236",
		"119.147.111.190",
		"119.147.111.198",
		"119.147.111.197",
		"119.147.111.147",
		"119.147.111.126",
		"119.147.111.211",
		"119.147.111.137",
		"119.147.111.123",
		"119.147.111.237",
		"119.147.111.243",
		"119.147.111.177",
		"119.147.111.206",
		"119.147.111.225",
		"119.147.111.235",
		"119.147.111.199",
		"119.147.111.181",
		"119.147.111.109",
		"119.147.111.213",
		"119.147.111.221",
		"119.147.111.182",
		"119.147.111.146",
		"119.147.111.195",
		"119.147.111.220",
		"119.147.111.140",
		"119.147.111.142",
		"119.147.111.113",
		"119.147.111.196",
		"119.147.111.183",
		"119.147.111.248",
		"119.147.111.136",
		"119.147.111.244",
		"119.147.111.239",
		"119.147.111.112",
		"119.147.111.128",
		"119.147.111.210",
		"119.147.111.218",
		"119.147.111.226",
		"119.147.111.228",
		"119.147.111.194",
		"119.147.111.224",
		"119.147.111.233",
		"119.147.111.200",
		"119.147.111.103",
	}

	// create a target list to scan
	targetCh, errCh := scan.CreateTargets(ipList, portList)
	targets := make([]plugins.Target, 0, 200)
	var allResults []plugins.Service // 存储所有扫描结果
	for {
		select {
		case target, ok := <-targetCh:
			if !ok {
				targetCh = nil //如果targetCh关闭，置nil以防止死锁
			} else {
				targets = append(targets, target) //添加目标到切片
				if len(targets) == 200 {          //如果目标达到200个
					results, err := scan.ScanTargetsMax(targets, fxConfig, 200)
					if err != nil {
						fmt.Println(err)
						continue
					}
					allResults = append(allResults, results...)
					targets = targets[:0] //清空目标，准备下一轮扫描
				}
			}

		case err, ok := <-errCh:
			if !ok {
				errCh = nil //如果errCh关闭，置nil以防止死锁
			}
			// 处理err,例如打印错误，或停止程序等
			fmt.Println(err)
		}

		if targetCh == nil && errCh == nil {
			if len(targets) > 0 { //如果还有剩余的目标没有扫描
				results, err := scan.ScanTargetsMax(targets, fxConfig, 200)
				if err != nil {
					fmt.Println(err)
				}
				allResults = append(allResults, results...)
			}
			break
		}
	}

	// process the results
	for _, result := range allResults {
		fmt.Printf("%s://%s:%d (%s/%s)\n", result.Protocol, result.Host, result.Port, result.Transport, result.Protocol)
		//fmt.Printf("%+v\n", result)

	}
}
