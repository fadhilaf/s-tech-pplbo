package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", nil)
}
