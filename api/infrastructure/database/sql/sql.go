package sql

import _ "embed"

var (
	// ----------------------------------------------------------------
	// regarding general
	// ----------------------------------------------------------------
	//go:embed general/insert_table_info.sql
	InsertTableInformation string
	//go:embed general/update_table_info.sql
	UpdateTableInformation string
	//go:embed general/check_id_in_table_info.sql
	CheckIdInTableInfo string

	// ----------------------------------------------------------------
	// regarding location
	// ----------------------------------------------------------------
	//go:embed location/find_phone_number_list_by_id.sql
	FindPhoneNumberListByID string
	//go:embed location/find_address_by_id.sql
	FindAddressByID string
	//go:embed location/find_company_by_id.sql
	FindCompanyByID string
	//go:embed location/find_warehouse_by_id.sql
	FindWarehouseByID string
	//go:embed location/find_factory_by_id.sql
	FindFactoryByID string
	//go:embed location/find_end_user_list_by_id.sql
	FindEndUserListByID string
	//go:embed location/insert_end_user_list.sql
	InsertEndUserList string
	//go:embed location/insert_bss_end_user_list_id.sql
	InsertBssEndUserListID string

	// ----------------------------------------------------------------
	// regarding machine
	// ----------------------------------------------------------------
	//go:embed machine/find_machine_by_id.sql
	FindMachineByID string
	//go:embed machine/find_all_machines.sql
	FindAllMachines string
	//go:embed machine/insert_machine.sql
	InsertMachine string
	//go:embed machine/update_machine.sql
	UpdateMachine string
	//go:embed machine/stop_using_machine.sql
	StopUsingMachine string
	//go:embed machine/find_machine_list_by_id.sql
	FindMachineListByID string
	//go:embed machine/insert_machine_list.sql
	InsertMachineList string
	//go:embed machine/insert_bss_machine_list_id.sql
	InsertBssMachineListID string

	// ----------------------------------------------------------------
	// regarding item
	// ----------------------------------------------------------------
	//go:embed item/find_all_items.sql
	FindAllItems string
	//go:embed item/find_all_items_in_detail.sql
	FindAllItemsInDetail string
	//go:embed item/find_item_by_id.sql
	FindItemByID string
	//go:embed item/find_item_in_detail_by_id.sql
	FindItemInDetailByID string
	//go:embed item/find_items_by_name.sql
	FindItemsByName string
	//go:embed item/insert_item.sql
	InsertItem string
	//go:embed item/update_item.sql
	UpdateItem string
	//go:embed item/stop_using_item.sql
	StopUsingItem string

	// ----------------------------------------------------------------
	// regarding item's category
	// ----------------------------------------------------------------
	//go:embed item/category/find_category_by_id.sql
	FindItemCategoryByID string
	//go:embed item/category/find_all_categories.sql
	FindAllItemCategories string
	//go:embed item/category/insert_item_category.sql
	InsertItemCategory string
	//go:embed item/category/update_item_category.sql
	UpdateItemCategory string
	//go:embed item/category/stop_using_item_category.sql
	StopUsingItemCategory string

	// ----------------------------------------------------------------
	// regarding item's status
	// ----------------------------------------------------------------
	//go:embed item/status/find_status_by_id.sql
	FindItemStatusByID string
	//go:embed item/status/find_all_statuses.sql
	FindAllItemStatuses string
	//go:embed item/status/insert_item_status.sql
	InsertItemStatus string
	//go:embed item/status/update_item_status.sql
	UpdateItemStatus string
	//go:embed item/status/stop_using_item_status.sql
	StopUsingItemStatus string

	// ----------------------------------------------------------------
	// regarding item's unit type
	// ----------------------------------------------------------------
	//go:embed item/unit/find_all_units.sql
	FindAllItemUnits string
	//go:embed item/unit/find_unit_by_id.sql
	FindItemUnitByID string
	//go:embed item/unit/insert_item_unit.sql
	InsertItemUnit string
	//go:embed item/unit/update_item_unit.sql
	UpdateItemUnit string
	//go:embed item/unit/stop_using_item_unit.sql
	StopUsingItemUnit string

	// ----------------------------------------------------------------
	// regarding inventory
	// ----------------------------------------------------------------
	//go:embed inventory/find_inventory.sql
	FindInventory string
	//go:embed inventory/find_all_inventories.sql
	FindAllInventories string
	//go:embed inventory/count_item_in_inventory.sql
	CountItemInInventory string
	//go:embed inventory/count_inventories.sql
	CountInventories string

	// ----------------------------------------------------------------
	// regarding process
	// ----------------------------------------------------------------
	//go:embed item/process/find_process_by_id.sql
	FindItemProcessByID string
	//go:embed item/process/find_all_processes.sql
	FindAllItemProcesses string
	//go:embed item/process/insert_item_process.sql
	InsertItemProcess string
	//go:embed item/process/update_item_process.sql
	UpdateItemProcess string
	//go:embed item/process/stop_using_item_process.sql
	StopUsingItemProcess string

	// ----------------------------------------------------------------
	// regarding user
	// ----------------------------------------------------------------
	//go:embed user/insert_user.sql
	InsertUser string
	//go:embed user/find_user_by_id.sql
	FindUserByID string
	//go:embed home/find_user_by_name.sql
	FindUserByName string
	//go:embed user/find_all_users.sql
	FindAllUsers string

	// ----------------------------------------------------------------
	// regarding production
	// ----------------------------------------------------------------
	//go:embed production/find_all_productions.sql
	FindAllProductions string
	//go:embed production/find_production_by_id.sql
	FindProductionByID string
	//go:embed production/find_production_by_item_id.sql
	FindProductionByItemID string
	//go:embed production/count_bss_consumption_list.sql
	CountBssConsumptionList string
	//go:embed production/find_item_in_inventory.sql
	FindItemInInventory string
	//go:embed production/insert_production.sql
	InsertProduction string
	//go:embed production/insert_consumption_list_id.sql
	InsertConsumptionListID string
	//go:embed production/insert_consumption_list.sql
	InsertConsumptionList string
	//go:embed production/upsert_inventory.sql
	UpsertInventory string
)

var (
	//go:embed machine/fake/fake_insert_machine.sql
	FakeInsertMachine string
)
