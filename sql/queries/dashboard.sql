-- name: GetClickTimeline :many
SELECT
    DATE(clicked_at) AS date,
    COUNT(*) AS clicks
FROM click_events
WHERE link_id = $1
GROUP BY DATE(clicked_at)
ORDER BY DATE(clicked_at);


-- name: GetDashboardSummary :one
SELECT
    COUNT(*) AS total_links,
    COALESCE(SUM(click_count),0) AS total_clicks,
    COUNT(*) FILTER (WHERE active = true) AS active_links
FROM links;