package inventory

import (
	"api/domain/model/general"
	"api/domain/model/item"
	"api/domain/model/location"
)

type Regulate struct {
	ID                    RegulateID                 `json:"regulateID"`
	No                    uint16                     `json:"regulateNo"`
	ItemID                item.ItemID                `json:"regulateItemID"`
	Item                  item.Item                  `json:"regulateItem"`
	ProcessID             item.ProcessID             `json:"regulateProcessID"`
	Process               item.Process               `json:"regulateProcess"`
	Lot                   string                     `json:"regulateLot"`
	Branch                string                     `json:"regulateBranch"`
	WarehouseIDBefore     location.WarehouseID       `json:"regulateWarehouseIDBefore"`
	WarehouseBefore       location.Warehouse         `json:"regulateWarehouseBefore"`
	WarehouseIDAfter      location.WarehouseID       `json:"regulateWarehouseIDAfter"`
	WarehouseAfter        location.Warehouse         `json:"regulateWarehouseAfter"`
	NonDefectiveQtyBefore float32                    `json:"regulateNonDefectiveQtyBefore"`
	NonDefectiveQtyAfter  float32                    `json:"regulateNonDefectiveQtyAfter"`
	DefectiveQtyBefore    float32                    `json:"regulateDefectiveQtyBefore"`
	DefectiveQtyAfter     float32                    `json:"regulateDefectiveQtyAfter"`
	SuspendedQtyBefore    float32                    `json:"regulateSuspendedQtyBefore"`
	SuspendedQtyAfter     float32                    `json:"regulateSuspendedQtyAfter"`
	Remark                string                     `json:"regulateRemark"`
	TableInformationID    general.TableInformationID `json:"regulateTableInformationID"`
	TableInformation      general.TableInformation   `json:"regulateTableInformation"`
}

type RegulateID string
