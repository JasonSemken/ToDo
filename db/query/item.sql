-- name: CreateItem :one
INSERT INTO items (
  item
) VALUES (
  $1
)
RETURNING *;

-- name: GetItem :one
SELECT * FROM items
WHERE id = $1 LIMIT 1;

-- name: GetItemForUpdate :one
SELECT * FROM items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListItem :many
SELECT * FROM items
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateItem :one
UPDATE items
SET item = $2
WHERE id = $1
RETURNING *;
