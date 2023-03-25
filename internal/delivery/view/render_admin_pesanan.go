package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderAdminPesanan(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_pesanan.gohtml", gin.H{})
}
