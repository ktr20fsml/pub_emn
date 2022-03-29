UPDATE
    emn.mst_item_status
SET
    status_name = :status_name,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_item_status_id = :mst_item_status_id
;
