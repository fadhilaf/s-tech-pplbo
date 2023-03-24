// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	ProductID   uuid.UUID
	Quantity    int32
	Status      interface{}
	Description string
	CreatedAt   time.Time
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int32
	Stock       int32
	IsService   bool
	Description string
	Image       string
	CreatedAt   time.Time
}

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	Address      string
	Phone        string
	CreatedAt    time.Time
}
