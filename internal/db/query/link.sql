-- name: LinkAdd :one
WITH inserted AS (
    INSERT INTO link (url, status)
        VALUES ($1, 'pending')
        ON CONFLICT (url) DO NOTHING
        RETURNING *
)
SELECT * FROM inserted
UNION ALL
SELECT * FROM link WHERE url = $1
LIMIT 1;

-- name: LinkGetByID :one
SELECT * FROM link WHERE id = $1;

-- name: LinkGetList :many
SELECT * FROM link
LIMIT $1
OFFSET $2;

-- name: LinkImageAddMultiple :copyfrom
INSERT INTO link_image (link_id, url)
VALUES ($1, $2);

-- name: LinkSetStatus :exec
UPDATE link SET status = $2 WHERE id = $1;
