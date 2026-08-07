-- name: GetUserLinkByID :one
SELECT *
FROM links
WHERE id = $1
AND user_id = $2
LIMIT 1;


-- name: UpdateLinkURL :one
UPDATE links
SET original_url = $1,
    updated_at = NOW()
WHERE id = $2
AND user_id = $3
RETURNING *;


-- name: DisableLink :one
UPDATE links
SET active = FALSE,
    updated_at = NOW()
WHERE id = $1
AND user_id = $2
RETURNING *;


-- name: EnableLink :one
UPDATE links
SET active = TRUE,
    updated_at = NOW()
WHERE id = $1
AND user_id = $2
RETURNING *;