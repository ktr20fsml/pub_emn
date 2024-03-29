INSERT INTO
    emn.trn_production
    (
        trn_production_id,
        trn_production_instruction_id,
        mst_item_id,
        mst_process_id,
        lot,
        branch,
        mst_machine_id,
        mst_operator_id,
        bss_consumption_list_id,
        non_defective_qty,
        defective_qty,
        suspended_qty,
        produced_at,
        information,
        remark,
        is_canceled,
        table_information_id
    )
VALUES
    (
        :trn_production_id,
        :trn_production_instruction_id,
        :mst_item_id,
        :mst_process_id,
        :lot,
        :branch,
        :mst_machine_id,
        :mst_operator_id,
        :bss_consumption_list_id,
        :non_defective_qty,
        :defective_qty,
        :suspended_qty,
        :produced_at,
        :information,
        :remark,
        :is_canceled,
        :table_information_id
    )
;