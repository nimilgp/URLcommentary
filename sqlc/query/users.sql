-- name: CreateUser :exec
INSERT INTO Users (
    UserName, FirstName, LastName, EmailId
) VALUES (
    $1, $2, $3, $4
);

-- name: UpdateUserName :exec
UPDATE Users
SET UserName = $1
WHERE UserId = $2;

-- name: UpdateEmailId :exec
UPDATE Users
SET EmailId = $1
WHERE UserId = $2;