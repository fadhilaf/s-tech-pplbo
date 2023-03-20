package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("user_id") == nil {
			c.AbortWithStatus(401)
		}
		c.Set("user_id", session.Get("user_id"))
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("is_admin") == nil {
			c.AbortWithStatus(401)
		}
		c.Set("is_admin", session.Get("is_admin"))
	}
}
