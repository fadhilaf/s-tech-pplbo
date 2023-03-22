package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderLogin(c *gin.Context) {
	message := utils.GetResponse(c)

	// we'll need this later
	userId := utils.GetUserIdFromSession(c)

	utils.DeleteResponse(c)
	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"Message": message,
		"User":    userId,
	})
}
