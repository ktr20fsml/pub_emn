UPDATE
    emn.mst_machine
SET
    machine_name = :machine_name,
    mst_factory_id = :mst_factory_id,
    maker_id = :maker_id,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_machine_id = :mst_machine_id
;
