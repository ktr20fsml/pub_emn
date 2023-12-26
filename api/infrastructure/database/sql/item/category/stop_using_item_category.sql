UPDATE
    emn.mst_item_category
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_item_category_id = :mst_item_category_id
;
