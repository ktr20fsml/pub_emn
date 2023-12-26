package machine

import (
	domainGeneral "api/domain/model/general"
	domainMachine "api/domain/model/machine"
	"api/infrastructure/dto/general"
)

type MachineList struct {
	ID                 string                   `db:"bss_machine_list_id"`
	No                 uint16                   `db:"machine_list_no"`
	MachineID          string                   `db:"mst_machine_id"`
	Machine            Machine                  `db:"machine"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToMachineListData(reqMachineList *domainMachine.MachineList) *MachineList {
	return &MachineList{
		ID:                 string(reqMachineList.ID),
		No:                 reqMachineList.No,
		MachineID:          string(reqMachineList.MachineID),
		Machine:            *ConvertToMachineData(&reqMachineList.Machine),
		TableInformationID: string(reqMachineList.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqMachineList.TableInformation),
	}
}
func ConvertToMachineListsDatas(reqMachineList []*domainMachine.MachineList) []*MachineList {
	machineList := make([]*MachineList, len(reqMachineList))

	for i, list := range reqMachineList {
		machineList[i] = ConvertToMachineListData(list)
	}

	return machineList
}

func ConvertToMachineListDomain(machineList *MachineList) *domainMachine.MachineList {
	return &domainMachine.MachineList{
		ID:                 domainMachine.MachineListID(machineList.ID),
		No:                 machineList.No,
		MachineID:          domainMachine.MachineID(machineList.MachineID),
		Machine:            *ConvertToMachineDomain(&machineList.Machine),
		TableInformationID: domainGeneral.TableInformationID(machineList.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&machineList.TableInformation),
	}
}
func ConvertToMachineListsDomains(machineList []*MachineList) []*domainMachine.MachineList {
	resMachineList := make([]*domainMachine.MachineList, len(machineList))

	for i, list := range machineList {
		resMachineList[i] = ConvertToMachineListDomain(list)
	}

	return resMachineList
}
