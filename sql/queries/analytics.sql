-- name: CreateClickEvent :exec
INSERT INTO click_events (
    link_id,
    ip_address,
    user_agent,
    referer
)
VALUES (
    $1,
    $2,
    $3,
    $4
);


-- name: GetLinkAnalytics :many
SELECT
    clicked_at,
    ip_address,
    user_agent,
    referer
FROM click_events
WHERE link_id = $1
ORDER BY clicked_at DESC;