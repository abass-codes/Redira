-- name: CreateUserLink :one
INSERT INTO links (
    user_id,
    original_url,
    short_code
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;


-- name: GetUserLinks :many
SELECT
    links.*,
    MAX(analytics_events.created_at) AS last_clicked_at
FROM links
LEFT JOIN analytics_events
ON links.id = analytics_events.link_id
WHERE links.user_id = $1
GROUP BY links.id
ORDER BY links.created_at DESC;


-- name: DeleteUserLink :exec
DELETE FROM links
WHERE id = $1
AND user_id = $2;