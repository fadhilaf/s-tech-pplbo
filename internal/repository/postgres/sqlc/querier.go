// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (sql.Result, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetOrderById(ctx context.Context, id uuid.UUID) (Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
	GetProduct(ctx context.Context) ([]Product, error)
	GetProductById(ctx context.Context, id uuid.UUID) (Product, error)
	GetProductByQuery(ctx context.Context, name string) ([]Product, error)
	GetUser(ctx context.Context) ([]GetUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (User, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (sql.Result, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (sql.Result, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
