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

// These import statements ensure that the init functions run in each plugin.
// When a new plugin is added, this list should be updated.

import (
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/dhcp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/dns"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/echo"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ftp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/http"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/imap"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ipmi"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ipsec"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/jdwp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/kafka/kafkaNew"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/kafka/kafkaOld"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ldap"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/linuxrpc"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/modbus"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/mqtt/mqtt3"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/mqtt/mqtt5"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/mssql"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/mysql"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/netbios"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ntp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/openvpn"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/oracledb"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/pop3"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/postgresql"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/rdp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/redis"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/rsync"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/rtsp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/smb"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/smtp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/snmp"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/ssh"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/stun"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/telnet"
	_ "gitlab.4399.cn/gz007/fingerprintx_pro/pkg/plugins/services/vnc"
)
