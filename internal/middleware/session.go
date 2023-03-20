package middleware

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(db *sql.DB) gin.HandlerFunc {
	store, err := postgres.NewStore(db, []byte("secret"))
	if err != nil {
		fmt.Println("Terjadi error dalam pembuatan session store", err)
	}

	return sessions.Sessions("mysession", store)
}
