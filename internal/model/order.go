package model

import (
	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID
	ProductID    uuid.UUID
	ProductName  string
	BuyerID      uuid.UUID
	BuyerName    string
	BuyerAddress string
	BuyerPhone   string
	Quantity     int32
	Status       string
	Description  string
}

type CreateOrderFormRequest struct {
	ProductID   uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity    int32     `json:"quantity" form:"quantity" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
}

type CreateOrderRequest struct {
	Form   CreateOrderFormRequest
	UserID uuid.UUID
}

type GetOrderByUserIdRequest struct {
	UserID uuid.UUID
}
