SELECT
    unit.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.mst_item_unit unit
JOIN
	emn.table_information as table_info
	ON
		unit.table_information_id = table_info.table_information_id
WHERE
    mst_item_unit_id = $1
;