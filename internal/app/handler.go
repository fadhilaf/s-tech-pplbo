package app

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/route"
	"github.com/gin-gonic/gin"
)

func (app *App) viewHandler(router *gin.RouterGroup) {
	viewGroup := router.Group("/")
	route.ViewRoutes(viewGroup, app.delivery.view)
}

func (app *App) apiHandler(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	route.UserRoutes(userGroup, app.delivery.user)
	authGroup := router.Group("/auth")
	route.AuthRoutes(authGroup, app.delivery.auth)
}
