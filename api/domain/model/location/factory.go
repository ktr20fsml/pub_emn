package location

import (
	"api/domain/model/general"
	"time"
)

type Factory struct {
	ID                 FactoryID                  `json:"factoryID,omitempty"`
	Name               string                     `json:"factoryName,omitempty"`
	CompanyID          CompanyID                  `json:"factoryCompanyID,omitempty"`
	Company            Company                    `json:"factoryCompany,omitempty"`
	AddressID          AddressID                  `json:"factoryAddressID,omitempty"`
	Address            Address                    `json:"factoryAddress,omitempty"`
	Remark             string                     `json:"factoryRemark,omitempty"`
	StopUsing          time.Time                  `json:"factoryStopUsing,omitempty"`
	TableInformationID general.TableInformationID `json:"factoryTableInformationID,omitempty"`
	TableInformation   general.TableInformation   `json:"factoryTableInformation,omitempty"`
}

type FactoryID string
