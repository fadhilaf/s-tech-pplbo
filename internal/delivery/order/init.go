package delivery

import (
	usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/order"
	"github.com/gin-gonic/gin"
)

type OrderDelivery interface {
	CreateOrder(ctx *gin.Context)
	DeliveredOrder(ctx *gin.Context)
	ProcessingOrder(ctx *gin.Context)
}

var _ OrderDelivery = &orderHandler{}

func NewOrderDelivery(usecase usecase.OrderUsecase) OrderDelivery {
	return &orderHandler{
		usecase: usecase,
	}
}

// type userHandler ini ditambahi satu per satu per file selain init
type orderHandler struct {
	usecase usecase.OrderUsecase
}
