package machine

import "api/domain/model/general"

type MachineList struct {
	ID                 MachineListID              `json:"machineListID,omitempty"`
	No                 uint16                     `json:"machineListNo"`
	MachineID          MachineID                  `json:"machineListMachineID"`
	Machine            Machine                    `json:"machineListMachine"`
	TableInformationID general.TableInformationID `json:"machineListTableInformationID"`
	TableInformation   general.TableInformation   `json:"machineListTableInformation"`
}

type MachineListID string
