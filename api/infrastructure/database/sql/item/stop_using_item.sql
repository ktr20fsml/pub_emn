UPDATE
    emn.mst_item
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_item_id = :mst_item_id
;
