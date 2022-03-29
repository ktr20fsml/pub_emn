SELECT
    company.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.mst_company company
JOIN
	emn.table_information as table_info
	ON
		company.table_information_id = table_info.table_information_id
WHERE
    mst_company_id = $1
;
