SELECT
    machine_list.*,
	table_info.created_at "table_information.created_at",
	table_info.created_by "table_information.created_by",
	table_info.updated_at "table_information.updated_at",
	table_info.updated_by "table_information.updated_by"
FROM
    emn.machine_list machine_list
JOIN
	emn.table_information as table_info
	ON
		machine_list.table_information_id = table_info.table_information_id
WHERE
    bss_machine_list_id = $1
;
