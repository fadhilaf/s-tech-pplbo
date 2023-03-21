package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderAdmin(c *gin.Context) {
	message := utils.GetResponse(c)

	if message != "" {
		utils.DeleteResponse(c)

		c.HTML(http.StatusOK, "admin.gohtml", gin.H{
			"Message": message,
		})
	} else {
		c.HTML(http.StatusOK, "admin.gohtml", gin.H{
			"Message": "",
		})
	}
}
