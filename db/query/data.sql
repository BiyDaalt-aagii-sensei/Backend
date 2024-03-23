-- name: ListGender :many
SELECT 
    *
FROM
    "data"
WHERE 
    "gender" = sqlc.arg('gender') :: VARCHAR;


-- name: GetGender :many

SELECT
    "gender"
FROM
    "data";

-- name: GetAge :many

SELECT
    "age"
FROM
    "data";

