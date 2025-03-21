// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package dblayer

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FullName, EmailId
) VALUES (
    $1, $2, $3
)
`

type CreateUserParams struct {
	Username string
	Fullname string
	Emailid  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Username, arg.Fullname, arg.Emailid)
	return err
}

const retrieveUserDetails = `-- name: RetrieveUserDetails :one
SELECT userid, username, fullname, emailid, joineddate, aboutme, passwordhash
FROM Users
WHERE EmailId = $1
`

func (q *Queries) RetrieveUserDetails(ctx context.Context, emailid string) (User, error) {
	row := q.db.QueryRow(ctx, retrieveUserDetails, emailid)
	var i User
	err := row.Scan(
		&i.Userid,
		&i.Username,
		&i.Fullname,
		&i.Emailid,
		&i.Joineddate,
		&i.Aboutme,
		&i.Passwordhash,
	)
	return i, err
}

const retrivePasswordHash = `-- name: RetrivePasswordHash :one
SELECT PasswordHash
FROM Users
WHERE EmailId = $1
`

func (q *Queries) RetrivePasswordHash(ctx context.Context, emailid string) (string, error) {
	row := q.db.QueryRow(ctx, retrivePasswordHash, emailid)
	var passwordhash string
	err := row.Scan(&passwordhash)
	return passwordhash, err
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
