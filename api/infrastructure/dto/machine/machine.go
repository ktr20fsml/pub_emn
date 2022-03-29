package machine

import (
	domainGeneral "api/domain/model/general"
	domainLocation "api/domain/model/location"
	domainMachine "api/domain/model/machine"
	"api/infrastructure/dto/general"
	"api/infrastructure/dto/location"
	"time"
)

type Machine struct {
	ID                 string                   `db:"mst_machine_id"`
	Name               string                   `db:"machine_name"`
	FactoryID          string                   `db:"mst_factory_id"`
	Factory            location.Factory         `db:"factory"`
	MakerID            string                   `db:"maker_id"`
	Maker              location.Company         `db:"maker"`
	Remark             string                   `db:"remark"`
	StopUsing          time.Time                `db:"stop_using"`
	TableInformationID string                   `db:"table_information_id"`
	TableInformation   general.TableInformation `db:"table_information"`
}

type MachineID string

func ConvertToMachineData(reqMachine *domainMachine.Machine) *Machine {
	return &Machine{
		ID:                 string(reqMachine.ID),
		Name:               reqMachine.Name,
		FactoryID:          string(reqMachine.FactoryID),
		Factory:            *location.ConvertToFactoryData(&reqMachine.Factory),
		MakerID:            string(reqMachine.MakerID),
		Maker:              *location.ConvertToCompanyData(&reqMachine.Maker),
		Remark:             reqMachine.Remark,
		StopUsing:          reqMachine.StopUsing,
		TableInformationID: string(reqMachine.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationData(&reqMachine.TableInformation),
	}
}

func ConvertToMachinesDatas(reqMachines []*domainMachine.Machine) []*Machine {
	locations := make([]*Machine, len(reqMachines))

	for i, reqMachine := range reqMachines {
		locations[i] = ConvertToMachineData(reqMachine)
	}

	return locations
}

func ConvertToMachineDomain(machine *Machine) *domainMachine.Machine {
	return &domainMachine.Machine{
		ID:                 domainMachine.MachineID(machine.ID),
		Name:               machine.Name,
		FactoryID:          domainLocation.FactoryID(machine.FactoryID),
		Factory:            *location.ConvertToFactoryDomain(&machine.Factory),
		MakerID:            domainLocation.CompanyID(machine.MakerID),
		Maker:              *location.ConvertToCompanyDomain(&machine.Maker),
		Remark:             machine.Remark,
		StopUsing:          machine.StopUsing,
		TableInformationID: domainGeneral.TableInformationID(machine.TableInformationID),
		TableInformation:   *general.ConvertToTableInformationDomain(&machine.TableInformation),
	}
}

func ConvertToMachinesDomains(machines []*Machine) []*domainMachine.Machine {
	resMachines := make([]*domainMachine.Machine, len(machines))

	for i, machine := range machines {
		resMachines[i] = ConvertToMachineDomain(machine)
	}

	return resMachines
}
