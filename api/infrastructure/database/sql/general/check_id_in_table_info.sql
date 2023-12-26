SELECT
    COUNT(table_information_id)
FROM
    emn.table_information
WHERE
    table_information_id = $1
;