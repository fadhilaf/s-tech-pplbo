package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SaveUser(c *gin.Context, id string) {
	session := sessions.Default(c)
	session.Set("user_id", id)
	session.Options(sessions.Options{
		HttpOnly: true,
	})
	session.Save()
}

func SaveAdmin(c *gin.Context, isAdmin bool) {
	session := sessions.Default(c)
	session.Set("is_admin", isAdmin)
	session.Options(sessions.Options{
		HttpOnly: true,
	})
	session.Save()
}
