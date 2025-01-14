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
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Emailid   string `json:"emailid"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
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
	Emailid string `json:"emailid"`
	Userid  int32  `json:"userid"`
}

func (q *Queries) UpdateEmailId(ctx context.Context, arg *UpdateEmailIdParams) error {
	_, err := q.db.ExecContext(ctx, updateEmailId, arg.Emailid, arg.Userid)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE Users
SET UserName = $1
WHERE UserId = $2
`

type UpdateUserNameParams struct {
	Username string `json:"username"`
	Userid   int32  `json:"userid"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg *UpdateUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserName, arg.Username, arg.Userid)
	return err
}
