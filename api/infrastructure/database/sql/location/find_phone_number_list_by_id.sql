SELECT
    phone_number_list.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.phone_number_list phone_number_list
JOIN
	emn.table_information as table_info
	ON
		phone_number_list.table_information_id = table_info.table_information_id
WHERE
    bss_phone_number_list_id = $1
;
