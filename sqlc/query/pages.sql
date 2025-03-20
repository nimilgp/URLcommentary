-- name: DoesPageExist :one
SELECT EXISTS (
    SELECT PageId
    FROM Pages
    WHERE PageURL = $1
);

-- name: CreatePage :exec
INSERT INTO Pages (
    PageURL
) VALUES (
    $1
);

-- name: RetrievePageDetails :one
SELECT 
    PageId,
    CommentsCount,
    CreatedAt,
    PageSummary,
    PageScore
FROM Pages
WHERE PageURL = $1;

-- name: UpdatePageSummaryDetails :exec
UPDATE Pages
SET PageSummary = $1,
    PageScore = $2
WHERE PageId = $3;

-- name: IncreaseCommentCount :exec
UPDATE Pages
SET CommentsCount = CommentsCount + 1
WHERE PageId = $1;