-- name: GetRedirectLink :one
SELECT *
FROM links
WHERE short_code = $1
LIMIT 1;