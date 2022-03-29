package production

import (
	domainItem "api/domain/model/item"
	domainLocation "api/domain/model/location"
	domainProduction "api/domain/model/production"
	"api/infrastructure/dto/item"
	"api/infrastructure/dto/location"
)

type ConsumptionList struct {
	ID              string             `db:"bss_consumption_list_id"`
	No              uint16             `db:"consumption_list_no"`
	WarehouseID     string             `db:"mst_warehouse_id"`
	Warehouse       location.Warehouse `db:"-"`
	ItemID          string             `db:"mst_item_id"`
	Item            item.Item          `db:"-"`
	ProcessID       string             `db:"mst_process_id"`
	Process         item.Process       `db:"-"`
	Lot             string             `db:"lot"`
	Branch          string             `db:"branch"`
	NonDefectiveQty float32            `db:"non_defective_qty"`
	DefectiveQty    float32            `db:"defective_qty"`
	SuspendedQty    float32            `db:"suspended_qty"`
	IsUsedUp        bool               `db:"is_used_up"`
	TransactionType string             `db:"transaction_type"`
}

func ConvertToConsumptionListData(reqConsumption *domainProduction.ConsumptionList) *ConsumptionList {
	return &ConsumptionList{
		ID:              string(reqConsumption.ID),
		No:              reqConsumption.No,
		WarehouseID:     string(reqConsumption.WarehouseID),
		Warehouse:       *location.ConvertToWarehouseData(&reqConsumption.Warehouse),
		ItemID:          string(reqConsumption.ItemID),
		Item:            *item.ConvertToItemData(&reqConsumption.Item),
		ProcessID:       string(reqConsumption.ProcessID),
		Process:         *item.ConvertToProcessData(&reqConsumption.Process),
		Lot:             reqConsumption.Lot,
		Branch:          reqConsumption.Branch,
		NonDefectiveQty: reqConsumption.NonDefectiveQty,
		DefectiveQty:    reqConsumption.DefectiveQty,
		SuspendedQty:    reqConsumption.SuspendedQty,
		IsUsedUp:        reqConsumption.IsUsedUp,
		TransactionType: reqConsumption.TransactionType,
	}
}

func ConvertToConsumptionListsDatas(reqConsumptions []*domainProduction.ConsumptionList) []*ConsumptionList {
	consumptions := make([]*ConsumptionList, len(reqConsumptions))

	for i, reqConsumption := range reqConsumptions {
		c := ConvertToConsumptionListData(reqConsumption)

		consumptions[i] = c
	}

	return consumptions
}

func ConvertToConsumptionListDomain(consumptionList *ConsumptionList) *domainProduction.ConsumptionList {
	return &domainProduction.ConsumptionList{
		ID:              domainProduction.ConsumptionListID(consumptionList.ID),
		No:              consumptionList.No,
		WarehouseID:     domainLocation.WarehouseID(consumptionList.WarehouseID),
		Warehouse:       *location.ConvertToWarehouseDomain(&consumptionList.Warehouse),
		ItemID:          domainItem.ItemID(consumptionList.ItemID),
		Item:            *item.ConvertToItemDomain(&consumptionList.Item),
		ProcessID:       domainItem.ProcessID(consumptionList.ProcessID),
		Process:         *item.ConvertToProcessDomain(&consumptionList.Process),
		Lot:             consumptionList.Lot,
		Branch:          consumptionList.Branch,
		NonDefectiveQty: consumptionList.NonDefectiveQty,
		DefectiveQty:    consumptionList.DefectiveQty,
		SuspendedQty:    consumptionList.SuspendedQty,
		IsUsedUp:        consumptionList.IsUsedUp,
		TransactionType: consumptionList.TransactionType,
	}
}

func ConvertToConsumptionListsDomains(consumptionLists []*ConsumptionList) []*domainProduction.ConsumptionList {
	resConsumptions := make([]*domainProduction.ConsumptionList, len(consumptionLists))

	for i, consumption := range consumptionLists {
		c := ConvertToConsumptionListDomain(consumption)

		resConsumptions[i] = c
	}

	return resConsumptions
}
