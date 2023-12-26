INSERT INTO
    emn.consumption_list
    (
        bss_consumption_list_id,
        consumption_list_no,
        mst_warehouse_id,
        mst_item_id,
        mst_process_id,
        lot,
        branch,
        non_defective_qty,
        defective_qty,
        suspended_qty,
        transaction_type
    )
VALUES
    (
        :bss_consumption_list_id,
        :consumption_list_no,
        :mst_warehouse_id,
        :mst_item_id,
        :mst_process_id,
        :lot,
        :branch,
        :non_defective_qty,
        :defective_qty,
        :suspended_qty,
        :transaction_type 
    )
;