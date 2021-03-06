//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           

package ipTable

import (
	_ "fmt"
	_ "net"
	"utils/logging"
)

/*
#cgo CFLAGS: -I../../../netfilter/libiptables/include -I../../../netfilter/iptables/include
#cgo LDFLAGS: -L../../../netfilter/libiptables/lib -lip4tc
#include "ipTable.h"
*/
import "C"

const (
	ALL_RULE_STR = "all"

	// Error Messages
	INSERTING_RULE_ERROR = "adding ip rule to iptables failed: "
	DELETING_RULE_ERROR  = "deleting rule failed: "
)

type SysdIpTableHandler struct {
	logger   *logging.Writer
	ruleInfo map[string]C.ipt_config_t
}
