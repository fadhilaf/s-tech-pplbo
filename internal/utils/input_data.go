package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
)

func SaveInputData(c *gin.Context, data model.InputData) {
	session := sessions.Default(c)
	session.Set("input_data", data)
	session.Save()
}
