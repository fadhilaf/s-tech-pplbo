// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: orders.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :execresult
INSERT INTO orders (
  user_id, product_id, quantity, status, description
) VALUES (
  $1, $2, $3, $4, $5
)
`

type CreateOrderParams struct {
	UserID      uuid.UUID
	ProductID   uuid.UUID
	Quantity    int32
	Status      interface{}
	Description string
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrder,
		arg.UserID,
		arg.ProductID,
		arg.Quantity,
		arg.Status,
		arg.Description,
	)
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOrderById = `-- name: GetOrderById :one
SELECT id, user_id, product_id, quantity, status, description, created_at FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrderById(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderById, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.Quantity,
		&i.Status,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getOrders = `-- name: GetOrders :many
SELECT id, user_id, product_id, quantity, status, description, created_at FROM orders 
ORDER BY created_at DESC
`

func (q *Queries) GetOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Quantity,
			&i.Status,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
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

const getOrdersByUserId = `-- name: GetOrdersByUserId :many
SELECT id, user_id, product_id, quantity, status, description, created_at FROM orders
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetOrdersByUserId(ctx context.Context, userID uuid.UUID) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Quantity,
			&i.Status,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
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

const updateOrder = `-- name: UpdateOrder :execresult
UPDATE orders SET
  status = $2,
  description = $3
WHERE id = $1
`

type UpdateOrderParams struct {
	ID          uuid.UUID
	Status      interface{}
	Description string
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateOrder, arg.ID, arg.Status, arg.Description)
}

const updateOrderStatus = `-- name: UpdateOrderStatus :execresult
UPDATE orders SET
  status = $2
WHERE id = $1
`

type UpdateOrderStatusParams struct {
	ID     uuid.UUID
	Status interface{}
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateOrderStatus, arg.ID, arg.Status)
}
