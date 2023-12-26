SELECT
    COUNT(bss_consumption_list_id)
FROM
    emn.bss_consumption_list
WHERE
    bss_consumption_list_id = $1
;