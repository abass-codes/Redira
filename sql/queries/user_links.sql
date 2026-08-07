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
SELECT *
FROM links
WHERE user_id = $1
ORDER BY created_at DESC;


-- name: DeleteUserLink :exec
DELETE FROM links
WHERE id = $1
AND user_id = $2;