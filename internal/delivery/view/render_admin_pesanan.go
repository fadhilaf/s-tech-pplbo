package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderAdminPesanan(c *gin.Context) {
	message := utils.GetResponse(c)

	// Ambil data pesanan
	resOrder := handler.usecase.GetOrder()
	var orders []model.Order
	if resOrder.Status == http.StatusOK {
		orders, _ = resOrder.Data["orders"].([]model.Order)
	}

	c.HTML(http.StatusOK, "admin_pesanan.gohtml", gin.H{
		"Message": message,
		"Orders":  orders,
	})
}
