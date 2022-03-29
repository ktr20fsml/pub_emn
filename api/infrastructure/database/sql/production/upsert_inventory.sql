INSERT INTO 
    emn.inventory
VALUES
    (
        :mst_item_id,
        :mst_process_id,
        :lot,
        :branch,
        :mst_warehouse_id,
        :non_defective_qty,
        :defective_qty,
        :suspended_qty,
        :expiration_date,
        :is_used,
        :is_used_up
    )
ON CONFLICT ON CONSTRAINT inventory_pkey
DO UPDATE SET
        mst_item_id          = :mst_item_id,
        mst_process_id       = :mst_process_id,
        lot                  = :lot,
        branch               = :branch,
        mst_warehouse_id     = :mst_warehouse_id,
        non_defective_qty    = :non_defective_qty,
        defective_qty        = :defective_qty,
        suspended_qty        = :suspended_qty,
        expiration_date      = :expiration_date,
        is_used              = :is_used,
        is_used_up           = :is_used_up
;