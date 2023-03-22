package delivery

import (
	"net/http"

	// "github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderHome(c *gin.Context) {
	// userId := utils.GetUserIdFromSession(c)

	c.HTML(http.StatusOK, "index.gohtml", gin.H{})
}
