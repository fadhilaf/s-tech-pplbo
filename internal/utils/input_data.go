package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
)

func SaveAuthInputData(c *gin.Context, data model.AuthInputData) {
	session := sessions.Default(c)
	session.Set("auth_input_data", data)
	session.Save()
}
