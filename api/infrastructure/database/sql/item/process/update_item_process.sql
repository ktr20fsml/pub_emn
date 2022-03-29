UPDATE
    emn.mst_process
SET
    process_name = :process_name,
    mst_factory_id = :mst_factory_id,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_process_id = :mst_process_id
;
