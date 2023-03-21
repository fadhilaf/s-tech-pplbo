package route

import (
	delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, handler delivery.AuthDelivery) {
	router.POST("/login", handler.UserLogin)
	router.POST("/admin/login", handler.AdminLogin)
	router.POST("/logout", handler.Logout)
}
