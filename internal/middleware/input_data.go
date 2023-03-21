package middleware

import (
	"fmt"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckInputData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		inputDataSession := session.Get("input_data")

		inputData, ok := inputDataSession.(model.InputData)

		if ok {
			fmt.Println("input data is not nil")
			c.Set("input_data", inputData)
		} else {
			fmt.Println("input data is nil")
		}

	}
}
