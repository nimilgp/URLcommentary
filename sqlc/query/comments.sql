-- name: CreateComment :exec
INSERT INTO Comments (
    PageId,UserId,CommentData,ParentId
) VALUES (
    $1, $2, $3, $4
);

-- name: UpdateComment :exec
UPDATE Comments
SET CommentData = $1, EditedBool = True
WHERE CommentId = $2;

-- name: RetrieveOldestComments :many
SELECT CommentId, UserName, CreatedAt, EditedBool, CommentData
FROM Comments, Users
WHERE Comments.UserId = Users.UserId AND PageId = $1
ORDER BY CreatedAt ASC
LIMIT $2;

-- name: RetrieveNewstComments :many
SELECT CommentId, UserName, CreatedAt, EditedBool, CommentData
FROM Comments, Users
WHERE Comments.UserId = Users.UserId AND PageId = $1
ORDER BY CreatedAt DESC
LIMIT $2;

-- name: RetrieveSubComments :many
SELECT CommentId, UserName, CreatedAt, EditedBool, CommentData
FROM Comments, Users
WHERE Comments.UserId = Users.UserId  AND ParentId = $1
ORDER BY CreatedAt DESC
LIMIT $2;