-- name: GetRedirectLink :one
SELECT *
FROM links
WHERE LOWER(short_code) = LOWER($1)
LIMIT 1;