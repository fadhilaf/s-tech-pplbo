package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderDashboard(c *gin.Context) {
	message := utils.GetResponse(c)

	// Ambil query search
	search := c.Query("search")

	// Ambil data produk
	resProduct := handler.usecase.GetProductByKeyword(model.GetProductByKeywordRequest{Keyword: search})
	var products []model.Product
	if resProduct.Status == http.StatusOK {
		products, _ = resProduct.Data["products"].([]model.Product)
	}

	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"Message":  message,
		"Products": products,
		"Amount":   len(products),
		"Search":   search,
	})
}
