-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FullName, EmailId, PasswordHash
) VALUES (
    $1, $2, $3, $4
);

-- name: RetrivePasswordHash :one
SELECT PasswordHash
FROM Users
WHERE EmailId = $1;

-- name: RetrieveUserDetails :one
SELECT UserName, FullName, JoinedDate, AboutMe
FROM Users
WHERE UserId = $1;

-- name: UpdateUserDetails :exec
UPDATE Users
SET 
    UserName = $1,
    FullName = $2,
    AboutMe = $3
WHERE UserId = $4;

-- name: DoesUserExist :one
SELECT EXISTS (
    SELECT UserId
    FROM Users
    WHERE EmailId = $1
);