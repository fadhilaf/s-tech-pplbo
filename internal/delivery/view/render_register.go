package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderRegister(c *gin.Context) {
	//gaya REST API
	inputDataSession, _ := c.Get("inputData")

	inputData, ok := inputDataSession.(model.InputData)

	if ok {
		c.HTML(http.StatusOK, "register.gohtml", gin.H{
			"message": inputData.Message,
		})
	} else {
		c.HTML(http.StatusOK, "register.gohtml", gin.H{
			"message": "",
		})
	}
}
