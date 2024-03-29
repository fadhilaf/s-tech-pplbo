package route

import (
	delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/order"
	"github.com/FadhilAF/s-tech-pplbo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup, handler delivery.OrderDelivery) {
	userRoutes := router.Group("/", middleware.ShouldUser())
	userRoutes.POST("/", handler.CreateOrder)
	userRoutes.POST("/delivered", handler.DeliveredOrder)

	adminRoutes := router.Group("/", middleware.ShouldAdmin())
	adminRoutes.POST("/processing", handler.ProcessingOrder)
}
