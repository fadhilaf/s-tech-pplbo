package route

import (
	delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/view"
	"github.com/FadhilAF/s-tech-pplbo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ViewRoutes(router *gin.RouterGroup, handler delivery.ViewDelivery) {
	router.GET("/", handler.RenderHome)
	router.GET("/register", handler.RenderRegister)
	router.GET("/login", handler.RenderLogin)

	// router.GET("/blabla", middleware1, middleware2) // biso ck ini jg dibuatny

	// Perlu dibuat Routes baru

	// userGroup := router.Group("/user", middleware.CheckUser())
	// userGroup.GET("/order", handler.RenderOrder)

	adminGroup := router.Group("/admin", middleware.CheckAdmin())
	adminGroup.GET("/admin", handler.RenderAdmin)
	// adminGroup.GET("/admin/tambah", handler.RenderTambah)
	// adminGroup.GET("/admin/order", handler.RenderOrderAdmin)
}