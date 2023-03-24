package route

import (
	delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup, handler delivery.ProductDelivery) {
	router.POST("/", handler.CreateProduct)
}
