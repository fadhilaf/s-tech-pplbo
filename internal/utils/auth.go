package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUserToSession(c *gin.Context, id uuid.UUID) {
	session := sessions.Default(c)
	session.Set("user_id", id)
	session.Options(sessions.Options{
		HttpOnly: true,
	})
	session.Save()
}

func SaveAdminToSession(c *gin.Context, isAdmin bool) {
	session := sessions.Default(c)
	session.Set("is_admin", isAdmin)
	session.Options(sessions.Options{
		HttpOnly: true,
	})
	session.Save()
}

func RemoveAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user") != nil {
		session.Delete("user")
	}

	if session.Get("is_admin") != nil {
		session.Delete("is_admin")
	}
	session.Save()
}
