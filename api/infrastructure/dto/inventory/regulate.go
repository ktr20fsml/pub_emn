package inventory

import (
	domainGeneral "api/domain/model/general"
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	domainLocation "api/domain/model/location"
	"api/infrastructure/dto/general"
	"api/infrastructure/dto/item"
	"api/infrastructure/dto/location"
)

type Regulate struct {
	ID                    string                   `db:"bss_regulate_inventory_id"`
	No                    uint16                   `db:"regulate_inventory_no"`
	ItemID                string                   `db:"mst_item_id"`
	Item                  item.Item                `db:"item"`
	ProcessID             string                   `db:"mst_process_id"`
	Process               item.Process             `db:"process"`
	Lot                   string                   `db:"lot"`
	Branch                string                   `db:"branch"`
	WarehouseIDBefore     string                   `db:"mst_warehouse_id_before"`
	WarehouseBefore       location.Warehouse       `db:"warehouse_before"`
	WarehouseIDAfter      string                   `db:"mst_warehouse_id_after"`
	WarehouseAfter        location.Warehouse       `db:"warehouse_after"`
	NonDefectiveQtyBefore float32                  `db:"nondefective_qty_before"`
	NonDefectiveQtyAfter  float32                  `db:"nondefective_qty_after"`
	DefectiveQtyBefore    float32                  `db:"defective_qty_before"`
	DefectiveQtyAfter     float32                  `db:"defective_qty_after"`
	SuspendedQtyBefore    float32                  `db:"suspended_qty_before"`
	SuspendedQtyAfter     float32                  `db:"suspended_qty_after"`
	Remark                string                   `db:"remark"`
	TableInformationID    string                   `db:"table_information_id"`
	TableInformation      general.TableInformation `db:"table_information"`
}

func ConvertToRegulatesDatas(reqRegulates []*domainInventory.Regulate) []*Regulate {
	regulates := make([]*Regulate, len(reqRegulates))

	for i, regulate := range reqRegulates {
		regulates[i] = &Regulate{
			ID:                    string(regulate.ID),
			No:                    regulate.No,
			ItemID:                string(regulate.ItemID),
			Item:                  *item.ConvertToItemData(&regulate.Item),
			ProcessID:             string(regulate.ProcessID),
			Process:               *item.ConvertToProcessData(&regulate.Process),
			Lot:                   regulate.Lot,
			Branch:                regulate.Branch,
			WarehouseIDBefore:     string(regulate.WarehouseIDBefore),
			WarehouseBefore:       *location.ConvertToWarehouseData(&regulate.WarehouseBefore),
			WarehouseIDAfter:      string(regulate.WarehouseIDAfter),
			WarehouseAfter:        *location.ConvertToWarehouseData(&regulate.WarehouseAfter),
			NonDefectiveQtyBefore: regulate.NonDefectiveQtyBefore,
			NonDefectiveQtyAfter:  regulate.NonDefectiveQtyAfter,
			DefectiveQtyBefore:    regulate.DefectiveQtyBefore,
			DefectiveQtyAfter:     regulate.DefectiveQtyAfter,
			SuspendedQtyBefore:    regulate.SuspendedQtyBefore,
			SuspendedQtyAfter:     regulate.SuspendedQtyAfter,
			Remark:                regulate.Remark,
			TableInformationID:    string(regulate.Item.TableInformationID),
			TableInformation:      *general.ConvertToTableInformationData(&regulate.TableInformation),
		}
	}

	return regulates
}

func ConvertToRegulatesDomains(regulates []*Regulate) []*domainInventory.Regulate {
	resRegulates := make([]*domainInventory.Regulate, len(regulates))

	for i, regulate := range regulates {
		resRegulates[i] = &domainInventory.Regulate{
			ID:                    domainInventory.RegulateID(regulate.ID),
			No:                    regulate.No,
			ItemID:                domainItem.ItemID(regulate.ItemID),
			Item:                  *item.ConvertToItemDomain(&regulate.Item),
			ProcessID:             domainItem.ProcessID(regulate.ProcessID),
			Process:               *item.ConvertToProcessDomain(&regulate.Process),
			Lot:                   regulate.Lot,
			Branch:                regulate.Lot,
			WarehouseIDBefore:     domainLocation.WarehouseID(regulate.WarehouseIDBefore),
			WarehouseBefore:       *location.ConvertToWarehouseDomain(&regulate.WarehouseBefore),
			WarehouseIDAfter:      domainLocation.WarehouseID(regulate.WarehouseIDAfter),
			WarehouseAfter:        *location.ConvertToWarehouseDomain(&regulate.WarehouseAfter),
			NonDefectiveQtyBefore: regulate.NonDefectiveQtyBefore,
			NonDefectiveQtyAfter:  regulate.NonDefectiveQtyAfter,
			DefectiveQtyBefore:    regulate.DefectiveQtyBefore,
			DefectiveQtyAfter:     regulate.DefectiveQtyAfter,
			SuspendedQtyBefore:    regulate.SuspendedQtyBefore,
			SuspendedQtyAfter:     regulate.SuspendedQtyAfter,
			Remark:                regulate.Remark,
			TableInformationID:    domainGeneral.TableInformationID(regulate.TableInformationID),
			TableInformation:      *general.ConvertToTableInformationDomain(&regulate.TableInformation),
		}
	}

	return resRegulates
}
