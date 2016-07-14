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

package server

import (
	//"fmt"
	"strings"
)

type AlarmState struct {
	OwnerId        int
	EventId        int
	OwnerName      string
	EventName      string
	SrcObjName     string
	Description    string
	OccuranceTime  string
	ResolutionTime string
	SrcObjKey      string
	Severity       string
}

func (server *FMGRServer) GetBulkAlarm(fromIdx int, cnt int) (int, int, []AlarmState) {
	server.logger.Info("Inside Get Bulk alarm function ...")
	var nextIdx int
	var count int

	alarms := server.AlarmList.GetListOfEntriesFromRingBuffer()
	faults := server.FaultList.GetListOfEntriesFromRingBuffer()
	length := len(alarms)
	result := make([]AlarmState, cnt)
	var i int
	var j int
	for i, j = 0, fromIdx; i < cnt && j < length; j++ {
		alarmIntf := alarms[length-j-1]
		alarm := alarmIntf.(AlarmDatabaseKey)
		faultIntf := faults[alarm.FaultIdx]
		fault := faultIntf.(FaultDatabaseKey)
		result[i].OwnerId = fault.FaultId.DaemonId
		result[i].EventId = fault.FaultId.EventId
		result[i].OwnerName = fault.OwnerName
		result[i].EventName = fault.EventName
		str := strings.Split(string(fault.FObjKey), " ")
		result[i].SrcObjName = str[0]
		result[i].Description = fault.Description
		str = strings.Split(str[1], "[")
		str = strings.Split(str[1], "]")
		result[i].OccuranceTime = fault.OccuranceTime.String()
		result[i].SrcObjKey = str[0]
		if fault.Resolved == false {
			result[i].ResolutionTime = "N/A"
		} else {
			result[i].ResolutionTime = fault.ResolutionTime.String()
		}
		flt := server.FaultEventMap[alarm.FaultId]
		result[i].Severity = flt.AlarmSeverity
		i++
	}
	if j == length {
		nextIdx = 0
	}

	count = i

	return nextIdx, count, result
}
