SELECT
    *
FROM
    emn.inventory
WHERE
    mst_item_id = $1
AND
    mst_process_id = $2
AND
    lot = $3
AND
    branch = $4
;