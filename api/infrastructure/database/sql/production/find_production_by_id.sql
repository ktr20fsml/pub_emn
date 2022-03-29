SELECT
    production.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.trn_production production
JOIN
	emn.table_information as table_info
	ON
		production.table_information_id = table_info.table_information_id
WHERE
    trn_production_id = $1
;
