package utils

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUserToSession(c *gin.Context, id uuid.UUID) {
	session := sessions.Default(c)

	idString := id.String()

	session.Set("user_id", idString)

	session.Options(sessions.Options{
		// dikasi "/" biar artinyo accessible di seluruh route didalem "/" (artiny segalo route), defaultny dio di /auth/login gr gr kito buatny pas lagi di route itu
		Path:     "/",
		HttpOnly: true,
		MaxAge:   24 * 3600,
	})

	err := session.Save()
	if err != nil {
		fmt.Println("error session save:", err)
	}

	// fmt.Println("user id yg dimasukkan adalah:", idString)
}

func GetUserIdFromSession(c *gin.Context) uuid.UUID {
	session := sessions.Default(c)

	userId, ok := session.Get("user_id").(string)
	if ok {
		parsedUserId, err := uuid.Parse(userId)
		if err == nil {
			// fmt.Println("tersimpan user id:", userId)
			return parsedUserId
		}
	}
	return uuid.Nil
}

func SaveAdminToSession(c *gin.Context, isAdmin bool) {
	session := sessions.Default(c)
	session.Set("is_admin", isAdmin)
	session.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   24 * 3600,
	})
	err := session.Save()
	if err != nil {
		fmt.Println("error session save:", err)
	}
}

func GetAdminFromSession(c *gin.Context) bool {
	session := sessions.Default(c)

	isAdmin := session.Get("is_admin")

	if isAdmin == nil {
		return false
	} else {
		return isAdmin.(bool)
	}
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
