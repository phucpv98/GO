-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM `pre_go_acc_user_base`
WHERE user_account = ?;

-- name: CheckUserBaseExist :one
SELECT COUNT(*)
FROM pre_go_acc_user_base
WHERE user_account = ?;

-- name: AddUserBase :execresult
INSERT INTO pre_go_acc_user_base (
    user_account,
    user_password,
    user_salt,
    user_created_at,
    user_updated_at
) VALUES (
    ?, ?, ?, NOW(), NOW()
);