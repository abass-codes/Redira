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
SELECT
    id,
    original_url,
    short_code,
    title,
    click_count,
    expires_at,
    created_at,
    updated_at
FROM links
WHERE short_code = $1
LIMIT 1;

-- name: IncrementClickCount :exec
UPDATE links
SET
    click_count = click_count + 1,
    updated_at = NOW()
WHERE id = $1;