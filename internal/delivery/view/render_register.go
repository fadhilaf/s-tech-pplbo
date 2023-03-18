package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.gohtml", nil)
}
