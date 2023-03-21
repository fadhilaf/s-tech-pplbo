package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderTambah(c *gin.Context) {
	c.HTML(http.StatusOK, "tambah.gohtml", gin.H{})
}
