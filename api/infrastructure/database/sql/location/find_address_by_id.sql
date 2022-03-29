SELECT
    address.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.mst_address address
JOIN
	emn.table_information as table_info
	ON
		address.table_information_id = table_info.table_information_id
WHERE
    mst_address_id = $1
;
