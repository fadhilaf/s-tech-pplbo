package utils

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

// REST API
func ToWebServiceResponse(message string, status int, data gin.H) model.WebServiceResponse {
	return model.WebServiceResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

// MVC
func SaveResponse(c *gin.Context, message string) {
	c.SetCookie("response", message, 60, "/", "localhost", false, true)
}

func GetResponse(c *gin.Context) string {
	cookie, err := c.Cookie("response")

	if err != nil {
		return ""
	}

	return cookie
}

func DeleteResponse(c *gin.Context) {
	c.SetCookie("response", "", -1, "/", "localhost", false, true)
}
