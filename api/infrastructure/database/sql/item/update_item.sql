UPDATE
    emn.mst_item
SET
    item_number = :item_number,
    item_name = :item_name,
    abbreviation = :abbreviation,
    mst_item_category_id = :mst_item_category_id,
    mst_item_status_id = :mst_item_status_id,
    mst_item_unit_id = :mst_item_unit_id,
    bss_machine_list_id = :bss_machine_list_id,
    mst_warehouse_id = :mst_warehouse_id,
    lower_limit_inventory_qty = :lower_limit_inventory_qty,
    safety_inventory_qty = :safety_inventory_qty,
    upper_limit_inventory_qty = :upper_limit_inventory_qty,
    min_lot_qty = :min_lot_qty,
    max_lot_qty = :max_lot_qty,
    unit_price = :unit_price,
    validity_days = :validity_days,
    mst_client_id = :mst_client_id,
    mst_delivery_destination_id = :mst_delivery_destination_id,
    bss_end_user_list_id = :bss_end_user_list_id,
    rank = :rank,
    remark = :remark,
    stop_using = :stop_using,
    table_information_id = :table_information_id
WHERE
    mst_item_id = :mst_item_id
;
