//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package pluginManager

import (
	"errors"
	"fmt"
	"infra/platformd/objects"
	"infra/platformd/pluginManager/pluginCommon"
	"utils/logging"
)

type SysManager struct {
	logger logging.LoggerIntf
	plugin PluginIntf
}

var SysMgr SysManager

func (sMgr *SysManager) Init(logger logging.LoggerIntf, plugin PluginIntf) {
	sMgr.logger = logger
	sMgr.plugin = plugin
	sMgr.logger.Info("System Manager Init()")
}

func (sMgr *SysManager) Deinit() {
	sMgr.logger.Info("System Manager Deinit()")
}

func (sMgr *SysManager) GetPlatformSystemState(sysName string) (*objects.PlatformSystemState, error) {
	var retObj objects.PlatformSystemState
	var platInfo pluginCommon.PlatformSystemState
	if sMgr.plugin == nil {
		return nil, errors.New("Invalid platform plugin")
	}
	platInfo, err := sMgr.plugin.GetPlatformSystemState()
	if err != nil {
		sMgr.logger.Err(fmt.Sprintln("Error getting Platform Info"))
		return &retObj, err
	}

	retObj.ObjName = sysName
	retObj.ProductName = platInfo.ProductName
	retObj.SerialNum = platInfo.SerialNum
	retObj.Manufacturer = platInfo.Manufacturer
	retObj.Vendor = platInfo.Vendor
	retObj.Release = platInfo.Release
	retObj.PlatformName = platInfo.PlatformName
	retObj.ONIEVersion = platInfo.ONIEVersion

	return &retObj, nil
}
