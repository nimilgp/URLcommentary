// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package dblayer

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FullName, EmailId, PasswordHash
) VALUES (
    $1, $2, $3, $4
)
`

type CreateUserParams struct {
	Username     string
	Fullname     string
	Emailid      string
	Passwordhash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Username,
		arg.Fullname,
		arg.Emailid,
		arg.Passwordhash,
	)
	return err
}

const doesUserExist = `-- name: DoesUserExist :one
SELECT EXISTS (
    SELECT UserId
    FROM Users
    WHERE EmailId = $1
)
`

func (q *Queries) DoesUserExist(ctx context.Context, emailid string) (bool, error) {
	row := q.db.QueryRow(ctx, doesUserExist, emailid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const retrieveUserDetails = `-- name: RetrieveUserDetails :one
SELECT UserName, FullName, JoinedDate, AboutMe
FROM Users
WHERE UserId = $1
`

type RetrieveUserDetailsRow struct {
	Username   string
	Fullname   string
	Joineddate pgtype.Timestamp
	Aboutme    string
}

func (q *Queries) RetrieveUserDetails(ctx context.Context, userid int32) (RetrieveUserDetailsRow, error) {
	row := q.db.QueryRow(ctx, retrieveUserDetails, userid)
	var i RetrieveUserDetailsRow
	err := row.Scan(
		&i.Username,
		&i.Fullname,
		&i.Joineddate,
		&i.Aboutme,
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

const updateUserDetails = `-- name: UpdateUserDetails :exec
UPDATE Users
SET 
    UserName = $1,
    FullName = $2,
    AboutMe = $3
WHERE UserId = $4
`

type UpdateUserDetailsParams struct {
	Username string
	Fullname string
	Aboutme  string
	Userid   int32
}

func (q *Queries) UpdateUserDetails(ctx context.Context, arg UpdateUserDetailsParams) error {
	_, err := q.db.Exec(ctx, updateUserDetails,
		arg.Username,
		arg.Fullname,
		arg.Aboutme,
		arg.Userid,
	)
	return err
}
