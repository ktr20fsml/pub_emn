package item

import (
	"api/domain/model/general"
	"api/domain/model/location"
	"api/domain/model/machine"
	"time"
)

type Item struct {
	ID                     ItemID                     `json:"itemID"`
	Number                 string                     `json:"itemNumber"`
	Name                   string                     `json:"itemName"`
	Abbreviation           string                     `json:"itemAbbreviation"`
	CategoryID             CategoryID                 `json:"itemCategoryID"`
	Category               Category                   `json:"itemCategory"`
	StatusID               StatusID                   `json:"itemStatusID"`
	Status                 Status                     `json:"itemStatus"`
	UnitID                 UnitID                     `json:"itemUnitID"`
	Unit                   Unit                       `json:"itemUnit"`
	MachineListID          machine.MachineListID      `json:"itemMachineListID"`
	MachineList            []*machine.MachineList     `json:"itemMachineList"`
	WarehouseID            location.WarehouseID       `json:"itemWarehouseID"`
	Warehouse              location.Warehouse         `json:"itemWarehouse"`
	LowerLimitInventoryQty float32                    `json:"itemLowerLimitInventoryQty"`
	SafetyInventoryQty     float32                    `json:"itemSafetyInventoryQty"`
	UpperLimitInventoryQty float32                    `json:"itemUpperLimitInventoryQty"`
	MinLotQty              float32                    `json:"itemMinLotQty"`
	MaxLotQty              float32                    `json:"itemMaxLotQty"`
	UnitPrice              float32                    `json:"itemUnitPrice"`
	ValidityDays           uint32                     `json:"itemValidityDays"`
	ClientID               location.CompanyID         `json:"itemClientID"`
	Client                 location.Company           `json:"itemClient"`
	DeliveryDestinationID  location.CompanyID         `json:"itemDeliveryDestinationID"`
	DeliveryDestination    location.Company           `json:"itemDeliveryDestination"`
	EndUserListID          location.EndUserListID     `json:"itemEndUserListID"`
	EndUserList            []*location.EndUserList    `json:"itemEndUserList"`
	Rank                   string                     `json:"itemRank"`
	Remark                 string                     `json:"itemRemark"`
	StopUsing              time.Time                  `json:"itemStopUsing"`
	TableInformationID     general.TableInformationID `json:"itemTableInformationID"`
	TableInformation       general.TableInformation   `json:"itemTableInformation"`
}

type ItemID string
