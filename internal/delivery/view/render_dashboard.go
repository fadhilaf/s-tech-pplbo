package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderDashboard(c *gin.Context) {

	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{})
}
