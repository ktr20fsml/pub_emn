UPDATE
    emn.mst_machine
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_machine_id = :mst_machine_id
;
