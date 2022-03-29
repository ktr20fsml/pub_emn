SELECT
	item.*,
	category.mst_item_category_id "category.mst_item_category_id",
	category.category_name "category.category_name",
	category.remark "category.remark",
	category.stop_using "category.stop_using",
	category.table_information_id "category.table_information_id",
	status.mst_item_status_id "status.mst_item_status_id",
	status.status_name "status.status_name",
	status.remark "status.remark",
	status.stop_using "status.stop_using",
	status.table_information_id "status.table_information_id",
	unit.mst_item_unit_id "unit.mst_item_unit_id",
	unit.unit_name "unit.unit_name",
	unit.remark "unit.remark",
	unit.stop_using "unit.stop_using",
	unit.table_information_id "unit.table_information_id",
	warehouse.mst_warehouse_id "warehouse.mst_warehouse_id",
	warehouse.warehouse_name "warehouse.warehouse_name",
	warehouse.mst_company_id "warehouse.mst_company_id",
	warehouse.mst_address_id "warehouse.mst_address_id",
	warehouse.remark "warehouse.remark",
	warehouse.stop_using "warehouse.stop_using",
	warehouse.table_information_id "warehouse.table_information_id",
	client.mst_company_id "client.mst_company_id",
	client.company_name "client.company_name",
	client.mst_address_id "client.mst_address_id",
	client.remark "client.remark",
	client.stop_using "client.stop_using",
	client.table_information_id "client.table_information_id",
	delivery_dest.mst_company_id "delivery_destination.mst_company_id",
	delivery_dest.company_name "delivery_destination.company_name",
	delivery_dest.mst_address_id "delivery_destination.mst_address_id",
	delivery_dest.remark "delivery_destination.remark",
	delivery_dest.stop_using "delivery_destination.stop_using",
	delivery_dest.table_information_id "delivery_destination.table_information_id",
	table_info.table_information_id "table_information.table_information_id",
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.mst_item AS item
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
  	emn.mst_warehouse AS warehouse
  	ON
		item.mst_warehouse_id = warehouse.mst_warehouse_id
JOIN 
	emn.mst_company AS client
	ON
		item.mst_client_id = client.mst_company_id
JOIN 
	emn.mst_company AS delivery_dest
	ON
		item.mst_delivery_destination_id = delivery_dest.mst_company_id
JOIN
	emn.table_information as table_info
	ON
		item.table_information_id = table_info.table_information_id
WHERE
    item.item_name
    LIKE
        $1
;