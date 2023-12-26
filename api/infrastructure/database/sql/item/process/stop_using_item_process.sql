UPDATE
    emn.mst_process
SET
    stop_using = CAST(CURRENT_TIMESTAMP AS TIMESTAMP)
WHERE
    mst_process_id = :mst_process_id
;
