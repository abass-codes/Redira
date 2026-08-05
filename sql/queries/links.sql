-- name: CreateLink :one
INSERT INTO links (
    original_url,
    short_code
)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetLinkByShortCode :one
SELECT *
FROM links
WHERE short_code = $1
LIMIT 1;