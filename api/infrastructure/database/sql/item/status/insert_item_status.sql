INSERT INTO
    emn.mst_item_status
    (
        mst_item_status_id,
        status_name,
        remark,
        table_information_id
    )
VALUES
    (
        :mst_item_status_id,
        :status_name,
        :remark,
        :table_information_id
    )
;
