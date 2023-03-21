package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "order.gohtml", gin.H{})
}
