package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.gohtml", nil)
}
