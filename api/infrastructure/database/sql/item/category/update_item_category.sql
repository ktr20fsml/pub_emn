UPDATE
    emn.mst_item_category
SET
    category_name = :category_name,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_item_category_id = :mst_item_category_id
;
