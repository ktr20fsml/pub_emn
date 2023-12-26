SELECT
    mst_user_id,
    user_name,
    user_password,
    administrative
FROM
    emn.mst_user
WHERE
    user_name = $1
;