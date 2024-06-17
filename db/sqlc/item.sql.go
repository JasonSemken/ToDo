// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: item.sql

package db

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
  item
) VALUES (
  $1
)
RETURNING id, item
`

func (q *Queries) CreateItem(ctx context.Context, item string) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem, item)
	var i Item
	err := row.Scan(&i.ID, &i.Item)
	return i, err
}

const getItem = `-- name: GetItem :one
SELECT id, item FROM items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetItem(ctx context.Context, id int64) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(&i.ID, &i.Item)
	return i, err
}

const getItemForUpdate = `-- name: GetItemForUpdate :one
SELECT id, item FROM items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetItemForUpdate(ctx context.Context, id int64) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemForUpdate, id)
	var i Item
	err := row.Scan(&i.ID, &i.Item)
	return i, err
}

const listItem = `-- name: ListItem :many
SELECT id, item FROM items
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListItemParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListItem(ctx context.Context, arg ListItemParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItem, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Item); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET item = $2
WHERE id = $1
RETURNING id, item
`

type UpdateItemParams struct {
	ID   int64  `json:"id"`
	Item string `json:"item"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem, arg.ID, arg.Item)
	var i Item
	err := row.Scan(&i.ID, &i.Item)
	return i, err
}