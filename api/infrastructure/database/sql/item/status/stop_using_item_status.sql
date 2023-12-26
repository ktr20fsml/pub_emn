UPDATE
    emn.mst_item_status
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_item_status_id = :mst_item_status_id
;
