UPDATE
    emn.mst_item_unit
SET
    unit_name = :unit_name,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_item_unit_id = :mst_item_unit_id
;
