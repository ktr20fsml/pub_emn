package location

import (
	"api/domain/model/general"
	"time"
)

type Warehouse struct {
	ID                 WarehouseID                `json:"warehouseID,omitempty"`
	Name               string                     `json:"warehouseName,omitempty"`
	CompanyID          CompanyID                  `json:"warehouseCompanyID,omitempty"`
	Company            Company                    `json:"warehouseCompany,omitempty"`
	AddressID          AddressID                  `json:"warehouseAddressID,omitempty"`
	Address            Address                    `json:"warehouseAddress,omitempty"`
	Remark             string                     `json:"warehouseRemark,omitempty"`
	StopUsing          time.Time                  `json:"warehouseStopUsing,omitempty"`
	TableInformationID general.TableInformationID `json:"warehouseTableInformationID,omitempty"`
	TableInformation   general.TableInformation   `json:"warehouseTableInformation,omitempty"`
}

type WarehouseID string
