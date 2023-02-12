package model

import (
	"time"

  "database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password" db:"password_hash"`
  Name         string    `json:"name" db:"name"`
	Address      string    `json:"address" db:"address"`
	Phone        string    `json:"phone" db:"phone"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
  Name     string `json:"name" binding:"required"`
  Address sql.NullString `json:"address" binding:"required"`
  Phone sql.NullString `json:"phone" binding:"required"`
}

type GetUserRequest struct {
  ID uuid.UUID `json:"id" binding:"required"`
}

type UpdateUserRequest struct {
  Name     string `json:"name" binding:"required"`
  Address  string `json:"address" binding:"required"`
  Phone    string `json:"phone" binding:"required"`
}

type DeleteUserRequest struct {
  ID uuid.UUID `json:"id" binding:"required"`
}

