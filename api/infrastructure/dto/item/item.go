package item

import (
	domainGeneral "api/domain/model/general"
	domainItem "api/domain/model/item"
	domainLocation "api/domain/model/location"
	domainMachine "api/domain/model/machine"
	"api/infrastructure/dto/general"
	"api/infrastructure/dto/location"
	"api/infrastructure/dto/machine"
	"time"
)

type Item struct {
	ID                     string                   `db:"mst_item_id"`
	Number                 string                   `db:"item_number"`
	Name                   string                   `db:"item_name"`
	Abbreviation           string                   `db:"abbreviation"`
	CategoryID             string                   `db:"mst_item_category_id"`
	Category               Category                 `db:"category"`
	StatusID               string                   `db:"mst_item_status_id"`
	Status                 Status                   `db:"status"`
	UnitID                 string                   `db:"mst_item_unit_id"`
	Unit                   Unit                     `db:"unit"`
	MachineListID          string                   `db:"bss_machine_list_id"`
	MachineList            []*machine.MachineList   `db:""`
	WarehouseID            string                   `db:"mst_warehouse_id"`
	Warehouse              location.Warehouse       `db:"warehouse"`
	LowerLimitInventoryQty float32                  `db:"lower_limit_inventory_qty"`
	SafetyInventoryQty     float32                  `db:"safety_inventory_qty"`
	UpperLimitInventoryQty float32                  `db:"upper_limit_inventory_qty"`
	MinLotQty              float32                  `db:"min_lot_qty"`
	MaxLotQty              float32                  `db:"max_lot_qty"`
	UnitPrice              float32                  `db:"unit_price"`
	ValidityDays           uint32                   `db:"validity_days"`
	ClientID               string                   `db:"mst_client_id"`
	Client                 location.Company         `db:"client"`
	DeliveryDestinationID  string                   `db:"mst_delivery_destination_id"`
	DeliveryDestination    location.Company         `db:"delivery_destination"`
	EndUserListID          string                   `db:"bss_end_user_list_id"`
	EndUserList            []*location.EndUserList  `db:""`
	Rank                   string                   `db:"rank"`
	Remark                 string                   `db:"remark"`
	StopUsing              time.Time                `db:"stop_using"`
	TableInformationID     string                   `db:"table_information_id"`
	TableInformation       general.TableInformation `db:"table_information"`
}

func ConvertToItemData(reqItem *domainItem.Item) *Item {
	return &Item{
		ID:                     string(reqItem.ID),
		Number:                 reqItem.Number,
		Name:                   reqItem.Name,
		Abbreviation:           reqItem.Abbreviation,
		CategoryID:             string(reqItem.CategoryID),
		Category:               *ConvertToCategoryData(&reqItem.Category),
		StatusID:               string(reqItem.StatusID),
		Status:                 *ConvertToStatusData(&reqItem.Status),
		UnitID:                 string(reqItem.UnitID),
		Unit:                   *ConvertToUnitData(&reqItem.Unit),
		MachineListID:          string(reqItem.MachineListID),
		MachineList:            machine.ConvertToMachineListsDatas(reqItem.MachineList),
		WarehouseID:            string(reqItem.WarehouseID),
		Warehouse:              *location.ConvertToWarehouseData(&reqItem.Warehouse),
		LowerLimitInventoryQty: reqItem.LowerLimitInventoryQty,
		SafetyInventoryQty:     reqItem.SafetyInventoryQty,
		UpperLimitInventoryQty: reqItem.UpperLimitInventoryQty,
		MinLotQty:              reqItem.MinLotQty,
		MaxLotQty:              reqItem.MaxLotQty,
		UnitPrice:              reqItem.UnitPrice,
		ValidityDays:           reqItem.ValidityDays,
		ClientID:               string(reqItem.ClientID),
		Client:                 *location.ConvertToCompanyData(&reqItem.Client),
		DeliveryDestinationID:  string(reqItem.DeliveryDestinationID),
		DeliveryDestination:    *location.ConvertToCompanyData(&reqItem.DeliveryDestination),
		EndUserListID:          string(reqItem.EndUserListID),
		EndUserList:            location.ConvertToEndUserListsDatas(reqItem.EndUserList),
		Rank:                   reqItem.Rank,
		Remark:                 reqItem.Remark,
		StopUsing:              reqItem.StopUsing,
		TableInformationID:     string(reqItem.TableInformationID),
		TableInformation:       *general.ConvertToTableInformationData(&reqItem.TableInformation),
	}
}

func ConvertToItemsDatas(reqItems []*domainItem.Item) []*Item {
	items := make([]*Item, len(reqItems))

	for i, reqItem := range reqItems {
		items[i] = ConvertToItemData(reqItem)
	}

	return items
}

func ConvertToItemDomain(item *Item) *domainItem.Item {
	return &domainItem.Item{
		ID:                     domainItem.ItemID(item.ID),
		Number:                 item.Number,
		Name:                   item.Name,
		Abbreviation:           item.Abbreviation,
		CategoryID:             domainItem.CategoryID(item.CategoryID),
		Category:               *ConvertToCategoryDomain(&item.Category),
		StatusID:               domainItem.StatusID(item.StatusID),
		Status:                 *ConvertToStatusDomain(&item.Status),
		UnitID:                 domainItem.UnitID(item.UnitID),
		Unit:                   *ConvertToUnitDomain(&item.Unit),
		MachineListID:          domainMachine.MachineListID(item.MachineListID),
		MachineList:            machine.ConvertToMachineListsDomains(item.MachineList),
		WarehouseID:            domainLocation.WarehouseID(item.WarehouseID),
		Warehouse:              *location.ConvertToWarehouseDomain(&item.Warehouse),
		LowerLimitInventoryQty: item.LowerLimitInventoryQty,
		SafetyInventoryQty:     item.SafetyInventoryQty,
		UpperLimitInventoryQty: item.UpperLimitInventoryQty,
		MinLotQty:              item.MinLotQty,
		MaxLotQty:              item.MaxLotQty,
		UnitPrice:              item.UnitPrice,
		ValidityDays:           item.ValidityDays,
		ClientID:               domainLocation.CompanyID(item.ClientID),
		Client:                 *location.ConvertToCompanyDomain(&item.Client),
		DeliveryDestinationID:  domainLocation.CompanyID(item.DeliveryDestinationID),
		DeliveryDestination:    *location.ConvertToCompanyDomain(&item.DeliveryDestination),
		EndUserListID:          domainLocation.EndUserListID(item.EndUserListID),
		EndUserList:            location.ConvertToEndUserListsDomains(item.EndUserList),
		Rank:                   item.Rank,
		Remark:                 item.Remark,
		StopUsing:              item.StopUsing,
		TableInformationID:     domainGeneral.TableInformationID(item.TableInformationID),
		TableInformation:       *general.ConvertToTableInformationDomain(&item.TableInformation),
	}
}

func ConvertToItemsDomains(items []*Item) []*domainItem.Item {
	resItems := make([]*domainItem.Item, len(items))

	for i, item := range items {
		resItems[i] = ConvertToItemDomain(item)
	}

	return resItems
}
