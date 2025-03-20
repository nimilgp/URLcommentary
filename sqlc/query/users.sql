-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FullName, EmailId
) VALUES (
    $1, $2, $3
);

-- name: RetrivePasswordHash :one
SELECT PasswordHash
FROM Users
WHERE EmailId = $1;

-- name: RetrieveUserDetails :one
SELECT *
FROM Users
WHERE EmailId = $1;

-- name: UpdateUserName :exec
UPDATE Users
SET UserName = $1
WHERE UserId = $2;
