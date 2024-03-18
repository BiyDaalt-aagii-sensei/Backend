-- name: CreateUser :one
INSERT INTO "users" (
    "Username",
    "Password",
    "FirstName",
    "LastName",
    "Email"
)
VALUES (
    sqlc.arg('Username') :: VARCHAR,
    sqlc.arg('Password') :: VARCHAR,
    sqlc.arg('FirstName') :: VARCHAR,
    sqlc.arg('LastName') :: VARCHAR,
    sqlc.arg('Email') :: VARCHAR
) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    "users"
WHERE
    "Username" = sqlc.arg('Username') :: VARCHAR
LIMIT 1;


-- name: UpdatePasswordUser :one
UPDATE 
    "users"
SET 
    "Password" = sqlc.arg('Password')
WHERE 
    "Id" = sqlc.arg('Id')
RETURNING *;