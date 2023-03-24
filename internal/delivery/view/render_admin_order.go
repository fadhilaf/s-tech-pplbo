package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderAdminOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_order.gohtml", gin.H{})
}
