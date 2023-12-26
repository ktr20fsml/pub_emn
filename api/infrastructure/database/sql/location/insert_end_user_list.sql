INSERT INTO
    emn.end_user_list
    (
        bss_end_user_list_id,
        end_user_list_no,
        end_user_id,
        remark,
        table_information_id
    )
    VALUES
    (
        :bss_end_user_list_id,
        :end_user_list_no,
        :end_user_id,
        :remark,
        :table_information_id
    )
;
