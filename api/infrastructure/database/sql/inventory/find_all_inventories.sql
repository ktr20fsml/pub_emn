SELECT
    inventory.*,
    item.mst_item_id "item.mst_item_id",
    item.item_number "item.item_number",
    item.item_name "item.item_name",
    item.abbreviation "item.abbreviation",
    item.mst_item_category_id "item.mst_item_category_id",
    item.mst_item_status_id "item.mst_item_status_id",
    item.mst_item_unit_id "item.mst_item_unit_id",
    item.bss_machine_list_id "item.bss_machine_list_id",
    item.mst_warehouse_id "item.mst_warehouse_id",
    item.lower_limit_inventory_qty "item.lower_limit_inventory_qty",
    item.safety_inventory_qty "item.safety_inventory_qty",
    item.upper_limit_inventory_qty "item.upper_limit_inventory_qty",
    item.min_lot_qty "item.min_lot_qty",
    item.max_lot_qty "item.max_lot_qty",
    item.unit_price "item.unit_price",
    item.validity_days "item.validity_days",
    item.mst_client_id "item.mst_client_id",
    item.mst_delivery_destination_id "item.mst_delivery_destination_id",
    item.bss_end_user_list_id "item.bss_end_user_list_id",
    item.rank "item.rank",
    item.remark "item.remark",
    item.stop_using "item.stop_using",
    item.table_information_id "item.table_information_id",
    category.mst_item_category_id "item.category.mst_item_category_id",
    category.category_name "item.category.category_name",
    category.remark "item.category.remark",
    category.stop_using "item.category.stop_using",
    category.table_information_id "item.category.table_information_id",
    status.mst_item_status_id "item.status.mst_item_status_id",
    status.status_name "item.status.status_name",
    status.remark "item.status.remark",
    status.stop_using "item.status.stop_using",
    status.table_information_id "item.status.table_information_id",
    unit.mst_item_unit_id "item.unit.mst_item_unit_id",
    unit.unit_name "item.unit.unit_name",
    unit.remark "item.unit.remark",
    unit.stop_using "item.unit.stop_using",
    unit.table_information_id "item.unit.table_information_id",
    process.mst_process_id "process.mst_process_id",
    process.process_name "process.process_name",
    process.mst_factory_id "process.mst_factory_id",
    process.remark "process.remark",
    process.stop_using "process.stop_using",
    process.table_information_id "process.table_information_id",
    warehouse.mst_warehouse_id "warehouse.mst_warehouse_id",
    warehouse.warehouse_name "warehouse.warehouse_name",
    warehouse.mst_company_id "warehouse.mst_company_id",
    warehouse.mst_address_id "warehouse.mst_address_id",
    warehouse.remark "warehouse.remark",
    warehouse.stop_using "warehouse.stop_using",
    warehouse.table_information_id "warehouse.table_information_id"
FROM 
    emn.inventory AS inventory
JOIN
    emn.mst_process AS process
    ON
        inventory.mst_process_id = process.mst_process_id
JOIN
    emn.mst_warehouse AS warehouse
    ON
        inventory.mst_warehouse_id = warehouse.mst_warehouse_id
JOIN 
    emn.mst_item AS item
    ON
        inventory.mst_item_id = item.mst_item_id
JOIN
    emn.mst_item_category AS category
    ON
        item.mst_item_category_id = category.mst_item_category_id
JOIN
    emn.mst_item_status AS status
    ON
        item.mst_item_status_id = status.mst_item_status_id
JOIN
    emn.mst_item_unit AS unit
    ON
        item.mst_item_unit_id = unit.mst_item_unit_id
JOIN 
    emn.mst_company AS client
    ON
        item.mst_client_id = client.mst_company_id
JOIN 
    emn.mst_company AS delivery_dest
    ON
        item.mst_delivery_destination_id = delivery_dest.mst_company_id
;