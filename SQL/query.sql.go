// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package dbsqlc

import (
	"context"
)

const createGroup = `-- name: CreateGroup :exec
INSERT INTO groups (groupName)
VALUES ($1)
`

func (q *Queries) CreateGroup(ctx context.Context, groupname string) error {
	_, err := q.db.ExecContext(ctx, createGroup, groupname)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, password)
VALUES ($1, $2)
`

type CreateUserParams struct {
	Name     string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Password)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE name=$1
`

func (q *Queries) DeleteUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, name)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, password FROM users
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}

const getUserGroups = `-- name: GetUserGroups :many
SELECT
    groups.id AS group_id,
    groups.groupName AS group_name
FROM
    user_groups
        JOIN
    groups ON user_groups.group_id = groups.id
WHERE
    user_groups.user_id = $1
`

type GetUserGroupsRow struct {
	GroupID   int32
	GroupName string
}

func (q *Queries) GetUserGroups(ctx context.Context, userID int32) ([]GetUserGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserGroups, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserGroupsRow
	for rows.Next() {
		var i GetUserGroupsRow
		if err := rows.Scan(&i.GroupID, &i.GroupName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, password FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
