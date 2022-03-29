UPDATE
    emn.mst_item_unit
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_item_unit_id = :mst_item_unit_id
;
