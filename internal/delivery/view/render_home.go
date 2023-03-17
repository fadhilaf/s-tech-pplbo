package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{})
}
