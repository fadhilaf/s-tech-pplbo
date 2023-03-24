package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderRegister(c *gin.Context) {
	message := utils.GetResponse(c)

	c.HTML(http.StatusOK, "register.gohtml", gin.H{
		"Message": message,
	})
}
