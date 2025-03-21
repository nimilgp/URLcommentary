// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: comments.sql

package dblayer

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createChildComment = `-- name: CreateChildComment :one
INSERT INTO ChildComments (
    PageId,
    UserId,
    ParentCommentId,
    Content
) VALUES (
    $1, $2, $3, $4
)
RETURNING ChildCommentId
`

type CreateChildCommentParams struct {
	Pageid          int32
	Userid          int32
	Parentcommentid int32
	Content         string
}

func (q *Queries) CreateChildComment(ctx context.Context, arg CreateChildCommentParams) (int32, error) {
	row := q.db.QueryRow(ctx, createChildComment,
		arg.Pageid,
		arg.Userid,
		arg.Parentcommentid,
		arg.Content,
	)
	var childcommentid int32
	err := row.Scan(&childcommentid)
	return childcommentid, err
}

const createParentComment = `-- name: CreateParentComment :one
INSERT INTO ParentComments (
    PageId,
    UserId,
    Content
) VALUES (
    $1, $2, $3
)
RETURNING CommentId
`

type CreateParentCommentParams struct {
	Pageid  int32
	Userid  int32
	Content string
}

func (q *Queries) CreateParentComment(ctx context.Context, arg CreateParentCommentParams) (int32, error) {
	row := q.db.QueryRow(ctx, createParentComment, arg.Pageid, arg.Userid, arg.Content)
	var commentid int32
	err := row.Scan(&commentid)
	return commentid, err
}

const retrieveChildComments = `-- name: RetrieveChildComments :many
SELECT 
    UserName, 
    CreatedAt, 
    Content
FROM ChildComments, Users
WHERE PageId = $1 AND ParentCommentId = $2 AND ChildComments.UserId = Users.UserId  
LIMIT $2 OFFSET $3
`

type RetrieveChildCommentsParams struct {
	Pageid int32
	Limit  int32
	Offset int32
}

type RetrieveChildCommentsRow struct {
	Username  string
	Createdat pgtype.Timestamp
	Content   string
}

func (q *Queries) RetrieveChildComments(ctx context.Context, arg RetrieveChildCommentsParams) ([]RetrieveChildCommentsRow, error) {
	rows, err := q.db.Query(ctx, retrieveChildComments, arg.Pageid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RetrieveChildCommentsRow
	for rows.Next() {
		var i RetrieveChildCommentsRow
		if err := rows.Scan(&i.Username, &i.Createdat, &i.Content); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveNewestParentComments = `-- name: RetrieveNewestParentComments :many
SELECT 
    CommentId, 
    UserName, 
    CreatedAt, 
    Content, 
    SentimentScore, 
    ChildCommentCount
FROM ParentComments, Users
WHERE PageId = $1 AND ParentComments.UserId = Users.UserId
ORDER BY CreatedAt DESC
LIMIT $2 OFFSET $3
`

type RetrieveNewestParentCommentsParams struct {
	Pageid int32
	Limit  int32
	Offset int32
}

type RetrieveNewestParentCommentsRow struct {
	Commentid         int32
	Username          string
	Createdat         pgtype.Timestamp
	Content           string
	Sentimentscore    int32
	Childcommentcount int32
}

func (q *Queries) RetrieveNewestParentComments(ctx context.Context, arg RetrieveNewestParentCommentsParams) ([]RetrieveNewestParentCommentsRow, error) {
	rows, err := q.db.Query(ctx, retrieveNewestParentComments, arg.Pageid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RetrieveNewestParentCommentsRow
	for rows.Next() {
		var i RetrieveNewestParentCommentsRow
		if err := rows.Scan(
			&i.Commentid,
			&i.Username,
			&i.Createdat,
			&i.Content,
			&i.Sentimentscore,
			&i.Childcommentcount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveOldestParentComments = `-- name: RetrieveOldestParentComments :many
SELECT 
    CommentId, 
    UserName, 
    CreatedAt, 
    Content, 
    SentimentScore, 
    ChildCommentCount
FROM ParentComments, Users
WHERE PageId = $1 AND ParentComments.UserId = Users.UserId
ORDER BY CreatedAt ASC
LIMIT $2 OFFSET $3
`

type RetrieveOldestParentCommentsParams struct {
	Pageid int32
	Limit  int32
	Offset int32
}

type RetrieveOldestParentCommentsRow struct {
	Commentid         int32
	Username          string
	Createdat         pgtype.Timestamp
	Content           string
	Sentimentscore    int32
	Childcommentcount int32
}

func (q *Queries) RetrieveOldestParentComments(ctx context.Context, arg RetrieveOldestParentCommentsParams) ([]RetrieveOldestParentCommentsRow, error) {
	rows, err := q.db.Query(ctx, retrieveOldestParentComments, arg.Pageid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RetrieveOldestParentCommentsRow
	for rows.Next() {
		var i RetrieveOldestParentCommentsRow
		if err := rows.Scan(
			&i.Commentid,
			&i.Username,
			&i.Createdat,
			&i.Content,
			&i.Sentimentscore,
			&i.Childcommentcount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
