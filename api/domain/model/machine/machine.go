package machine

import (
	"api/domain/model/general"
	"api/domain/model/location"
	"time"
)

type Machine struct {
	ID                 MachineID                  `json:"machineID,omitempty"`
	Name               string                     `json:"machineName,omitempty"`
	FactoryID          location.FactoryID         `json:"machineFactoryID,omitempty"`
	Factory            location.Factory           `json:"machineFactory,omitempty"`
	MakerID            location.CompanyID         `json:"machineMakerID,omitempty"`
	Maker              location.Company           `json:"machineMaker,omitempty"`
	Remark             string                     `json:"machineRemark,omitempty"`
	StopUsing          time.Time                  `json:"machineStopUsing,omitempty"`
	TableInformationID general.TableInformationID `json:"machineTableInformationID,omitempty"`
	TableInformation   general.TableInformation   `json:"machineTableInformation,omitempty"`
}

type MachineID string
