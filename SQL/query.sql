-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, password)
VALUES ($1, $2);

-- name: DeleteUser :exec
DELETE FROM users WHERE name=$1;

-- name: CreateGroup :exec
INSERT INTO groups (groupName)
VALUES ($1);


-- name: GetUserGroups :many
SELECT
    groups.group_id AS group_id,
    groups.groupName AS group_name
FROM
    user_groups
        JOIN
    groups ON user_groups.group_id = groups.id
WHERE
    user_groups.user_id = $1;


-- get user explicit roles and implicit roles from groups
-- name: GetUserRoles :many


