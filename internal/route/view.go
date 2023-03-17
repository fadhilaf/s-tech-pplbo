package route

import (
	delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/view"

	"github.com/gin-gonic/gin"
)

func ViewRoutes(router *gin.RouterGroup, handler delivery.ViewDelivery) {
	router.GET("/", handler.RenderHome)
	// router.GET("/register", RenderRegister)
	// router.GET("/login", RenderLogin)
}
