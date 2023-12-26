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
	warehouse_address.mst_address_id "warehouse.address.mst_address_id",
	warehouse_address.building "warehouse.address.building",
	warehouse_address.street "warehouse.address.street",
	warehouse_address.city "warehouse.address.city",
	warehouse_address.state "warehouse.address.state",
	warehouse_address.province "warehouse.address.province",
	warehouse_address.region "warehouse.address.region",
	warehouse_address.prefecture "warehouse.address.prefecture",
	warehouse_address.zip "warehouse.address.zip",
	warehouse_address.postal_code "warehouse.address.postal_code",
	warehouse_address.country "warehouse.address.country",
	warehouse_address.bss_phone_number_list_id "warehouse.address.bss_phone_number_list_id",
	warehouse_address.remark "warehouse.address.remark",
	warehouse_address.stop_using "warehouse.address.stop_using",
	warehouse_address.table_information_id "warehouse.address.table_information_id",
	client.mst_company_id "client.mst_company_id",
	client.company_name "client.company_name",
	client.mst_address_id "client.mst_address_id",
	client.remark "client.remark",
	client.stop_using "client.stop_using",
	client.table_information_id "client.table_information_id",
	client_address.mst_address_id "client.address.mst_address_id",
	client_address.building "client.address.building",
	client_address.street "client.address.street",
	client_address.city "client.address.city",
	client_address.state "client.address.state",
	client_address.province "client.address.province",
	client_address.region "client.address.region",
	client_address.prefecture "client.address.prefecture",
	client_address.zip "client.address.zip",
	client_address.postal_code "client.address.postal_code",
	client_address.country "client.address.country",
	client_address.bss_phone_number_list_id "client.address.bss_phone_number_list_id",
	client_address.remark "client.address.remark",
	client_address.stop_using "client.address.stop_using",
	client_address.table_information_id "client.address.table_information_id",
	delivery_dest.mst_company_id "delivery_destination.mst_company_id",
	delivery_dest.company_name "delivery_destination.company_name",
	delivery_dest.mst_address_id "delivery_destination.mst_address_id",
	delivery_dest.remark "delivery_destination.remark",
	delivery_dest.stop_using "delivery_destination.stop_using",
	delivery_dest.table_information_id "delivery_destination.table_information_id",
	delivery_dest_address.mst_address_id "delivery_destination.address.mst_address_id",
	delivery_dest_address.building "delivery_destination.address.building",
	delivery_dest_address.street "delivery_destination.address.street",
	delivery_dest_address.city "delivery_destination.address.city",
	delivery_dest_address.state "delivery_destination.address.state",
	delivery_dest_address.province "delivery_destination.address.province",
	delivery_dest_address.region "delivery_destination.address.region",
	delivery_dest_address.prefecture "delivery_destination.address.prefecture",
	delivery_dest_address.zip "delivery_destination.address.zip",
	delivery_dest_address.postal_code "delivery_destination.address.postal_code",
	delivery_dest_address.country "delivery_destination.address.country",
	delivery_dest_address.bss_phone_number_list_id "delivery_destination.address.bss_phone_number_list_id",
	delivery_dest_address.remark "delivery_destination.address.remark",
	delivery_dest_address.stop_using "delivery_destination.address.stop_using",
	delivery_dest_address.table_information_id "delivery_destination.address.table_information_id",
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
  	emn.mst_address AS warehouse_address
  	ON
		warehouse.mst_address_id = warehouse_address.mst_address_id
JOIN 
	emn.mst_company AS client
	ON
		item.mst_client_id = client.mst_company_id
JOIN
  	emn.mst_address AS client_address
  	ON
		client.mst_address_id = client_address.mst_address_id
JOIN 
	emn.mst_company AS delivery_dest
	ON
		item.mst_delivery_destination_id = delivery_dest.mst_company_id
JOIN
  	emn.mst_address AS delivery_dest_address
  	ON
		delivery_dest.mst_address_id = delivery_dest_address.mst_address_id
JOIN
	emn.table_information as table_info
	ON
		item.table_information_id = table_info.table_information_id
;