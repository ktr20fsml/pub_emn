package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"time"
)

type Process struct {
	ID                 string                   `db:"mst_process_id"`
	Name               string                   `db:"process_name"`
	FactoryID          string                   `db:"mst_factory_id"`
	StopUsing          time.Time                `db:"stop_using"`
	Remark             string                   `db:"remark"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

func ConvertToProcessData(reqProcess *domainItem.Process) *Process {
	return &Process{
		ID:                 string(reqProcess.ID),
		Name:               reqProcess.Name,
		FactoryID:          string(reqProcess.FactoryID),
		StopUsing:          reqProcess.StopUsing,
		Remark:             reqProcess.Remark,
		TableInformationID: string(reqProcess.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqProcess.TableInformation),
	}

}

func ConvertToProcessesDatas(reqProcesses []*domainItem.Process) []*Process {
	processes := make([]*Process, len(reqProcesses))

	for i, reqProcess := range reqProcesses {
		processes[i] = ConvertToProcessData(reqProcess)
	}

	return processes
}

func ConvertToProcessDomain(process *Process) *domainItem.Process {
	return &domainItem.Process{
		ID:                 domainItem.ProcessID(process.ID),
		Name:               process.Name,
		FactoryID:          domainLocation.FactoryID(process.FactoryID),
		Remark:             process.Remark,
		StopUsing:          process.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(process.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&process.TableInformation),
	}
}

func ConvertToProcessesDomains(processes []*Process) []*domainItem.Process {
	resProcesses := make([]*domainItem.Process, len(processes))

	for i, process := range processes {
		resProcesses[i] = ConvertToProcessDomain(process)
	}

	return resProcesses
}
