-- name: CreateParentComment :one
INSERT INTO ParentComments (
    PageId,
    UserId,
    Content
) VALUES (
    $1, $2, $3
)
RETURNING CommentId;

-- name: CreateChildComment :one
INSERT INTO ChildComments (
    PageId,
    UserId,
    ParentCommentId,
    Content
) VALUES (
    $1, $2, $3, $4
)
RETURNING ChildCommentId;

-- name: RetrieveOldestParentComments :many
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
LIMIT $2 OFFSET $3;

-- name: RetrieveNewestParentComments :many
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
LIMIT $2 OFFSET $3;

-- name: RetrieveChildComments :many
SELECT 
    UserName, 
    CreatedAt, 
    Content
FROM ChildComments, Users
WHERE PageId = $1 AND ParentCommentId = $2 AND ChildComments.UserId = Users.UserId  
LIMIT $3 OFFSET $4;