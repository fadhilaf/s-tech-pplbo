package middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userId := session.Get("user_id")

		fmt.Println("tersimpan user id:", userId)

		if userId == nil {
			c.AbortWithStatus(401)
		}

		c.Set("user", userId)
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		isAdmin := session.Get("is_admin")

		if isAdmin == nil {
			c.AbortWithStatus(401)
		}
		c.Set("is_admin", isAdmin)
	}
}
