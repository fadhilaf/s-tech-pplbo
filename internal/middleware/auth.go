package middleware

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := utils.GetUserIdFromSession(c)

		if userId == uuid.Nil {
			c.AbortWithStatus(401)
		}

		c.Set("user", userId)
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		isAdmin := utils.GetAdminFromSession(c)

		if isAdmin == false {
			c.AbortWithStatus(401)
		}
		c.Set("is_admin", isAdmin)
	}
}
