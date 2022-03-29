SELECT
    COUNT(trn_production_id)
FROM
    emn.trn_production
WHERE
    trn_production_id = $1
;