// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package dblayer

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FirstName, LastName, EmailId
) VALUES (
    $1, $2, $3, $4
)
`

type CreateUserParams struct {
	Username  string
	Firstname string
	Lastname  string
	Emailid   string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Username,
		arg.Firstname,
		arg.Lastname,
		arg.Emailid,
	)
	return err
}

const updateEmailId = `-- name: UpdateEmailId :exec
UPDATE Users
SET EmailId = $1
WHERE UserId = $2
`

type UpdateEmailIdParams struct {
	Emailid string
	Userid  int32
}

func (q *Queries) UpdateEmailId(ctx context.Context, arg UpdateEmailIdParams) error {
	_, err := q.db.Exec(ctx, updateEmailId, arg.Emailid, arg.Userid)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE Users
SET UserName = $1
WHERE UserId = $2
`

type UpdateUserNameParams struct {
	Username string
	Userid   int32
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	_, err := q.db.Exec(ctx, updateUserName, arg.Username, arg.Userid)
	return err
}
