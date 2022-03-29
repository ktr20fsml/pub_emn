INSERT INTO
    emn.mst_process
    (
        mst_process_id,
        process_name,
        mst_factory_id,
        remark,
        table_information_id
    )
VALUES
    (
        :mst_process_id,
        :process_name,
        :mst_factory_id,
        :remark,
        :table_information_id
    )
;
