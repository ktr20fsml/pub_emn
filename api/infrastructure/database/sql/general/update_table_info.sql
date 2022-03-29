UPDATE
    emn.table_information
SET
    updated_at = :updated_at,
    updated_by = :updated_by
WHERE
    table_information_id = :table_information_id
;
