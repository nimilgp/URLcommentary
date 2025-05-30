-- name: DoesLikeExist :one
SELECT EXISTS (
    SELECT *
    FROM LikesHistory
    WHERE 
        PageId = $1 AND
        UserId = $2 AND
        CommentId = $3
);

-- name: CreateLikeHistory :exec
INSERT INTO LikesHistory (
    PageId,
    UserId,
    CommentId,
    LikeValue
) VALUES (
    $1, $2, $3, $4
);

-- name: RetrieveLikeHistory :many
SELECT CommentId, LikeValue
FROM LikesHistory
WHERE PageId = $1 AND UserId = $2;

-- name: RetrieveLike :one
SELECT LikeValue
FROM LikesHistory
WHERE PageId = $1 AND UserId = $2 AND CommentId = $3;

-- name: UpdateLikeHistory :exec
UPDATE LikesHistory
SET LikeValue = $1
WHERE 
    PageId = $2 AND
    UserId = $3 AND
    CommentId = $4;