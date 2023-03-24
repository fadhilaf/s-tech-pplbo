// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: products.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :execresult
INSERT INTO products (
  name, price, stock, is_service, description, image
) VALUES (
  $1, $2, $3, $4, $5, $6
)
`

type CreateProductParams struct {
	Name        string
	Price       int32
	Stock       int32
	IsService   bool
	Description string
	Image       string
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProduct,
		arg.Name,
		arg.Price,
		arg.Stock,
		arg.IsService,
		arg.Description,
		arg.Image,
	)
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :many
SELECT id, name, price, stock, is_service, description, image, created_at FROM products
ORDER BY name
`

func (q *Queries) GetProduct(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Stock,
			&i.IsService,
			&i.Description,
			&i.Image,
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

const getProductById = `-- name: GetProductById :one
SELECT id, name, price, stock, is_service, description, image, created_at FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductById(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Stock,
		&i.IsService,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const getProductByName = `-- name: GetProductByName :one
SELECT id, name, price, stock, is_service, description, image, created_at FROM products
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetProductByName(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByName, name)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Stock,
		&i.IsService,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const getProductByQuery = `-- name: GetProductByQuery :many
SELECT id, name, price, stock, is_service, description, image, created_at FROM products
WHERE name ILIKE $1
ORDER BY name
`

func (q *Queries) GetProductByQuery(ctx context.Context, name string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductByQuery, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Stock,
			&i.IsService,
			&i.Description,
			&i.Image,
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

const updateProduct = `-- name: UpdateProduct :execresult
UPDATE products SET
  name = $2,
  price = $3,
  stock = $4,
  description = $5,
  image = $6
WHERE id = $1
`

type UpdateProductParams struct {
	ID          uuid.UUID
	Name        string
	Price       int32
	Stock       int32
	Description string
	Image       string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.Stock,
		arg.Description,
		arg.Image,
	)
}
