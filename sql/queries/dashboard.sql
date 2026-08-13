-- name: GetClickTimeline :many
SELECT
    DATE(created_at) AS date,
    COUNT(*) AS clicks
FROM analytics_events
WHERE link_id = $1
GROUP BY DATE(created_at)
ORDER BY DATE(created_at);


-- name: GetDashboardSummary :one
SELECT
    COUNT(DISTINCT links.id) AS total_links,
    COUNT(analytics_events.id) AS total_clicks,
    COUNT(DISTINCT links.id) FILTER (WHERE links.active = true) AS active_links
FROM links
LEFT JOIN analytics_events
ON links.id = analytics_events.link_id;
