package delivery

import (
	usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

type ProductDelivery interface {
	CreateProduct(ctx *gin.Context)
}

var _ ProductDelivery = &productHandler{}

func NewProductDelivery(usecase usecase.ProductUsecase) ProductDelivery {
	return &productHandler{
		usecase: usecase,
	}
}

// type userHandler ini ditambahi satu per satu per file selain init
type productHandler struct {
	usecase usecase.ProductUsecase
}
