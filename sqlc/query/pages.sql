-- name: CreatePage :exec
INSERT INTO Pages (
    PageURL
) VALUES (
    $1
);

-- name: IncreaseCommentCount :exec
UPDATE Pages
SET CommentsCount = CommentsCount + 1
WHERE PageId = $1;

-- name: RetrievePageDetails :one
SELECT PageId, CommentsCount, CreatedAt
FROM Pages
WHERE PageURL = $1;