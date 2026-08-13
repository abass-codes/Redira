-- name: CreateClickEvent :exec
INSERT INTO analytics_events (
    link_id,
    ip_address,
    user_agent
)
VALUES (
    $1,
    $2,
    $3
);


-- name: GetLinkAnalytics :many
SELECT
    created_at,
    ip_address,
    user_agent
FROM analytics_events
WHERE link_id = $1
ORDER BY created_at DESC;
