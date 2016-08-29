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

package objects

type FaultState struct {
	OwnerId          int32
	EventId          int32
	OwnerName        string
	EventName        string
	SrcObjName       string
	Description      string
	OccuranceTime    string
	SrcObjKey        string
	SrcObjUUID       string
	ResolutionTime   string
	ResolutionReason string
}

type FaultStateGetInfo struct {
	EndIdx int
	Count  int
	More   bool
	List   []FaultState
}

const (
	ALL_EVENTS = "all"
)

type FaultEnable struct {
	OwnerName string
	EventName string
	Enable    bool
}

type FaultClear struct {
	OwnerName  string
	EventName  string
	SrcObjUUID string
}
