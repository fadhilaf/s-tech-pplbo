package delivery

import (
	_ "embed"

	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (handler *viewHandler) RenderDashboard(c *gin.Context) {
	var req model.GetProductByKeywordRequest

	message := utils.GetResponse(c)

	ok := utils.BindFormAndValidate(c, &req)
	if !ok {
		return
	}

	// Ambil data produk
	resProduct := handler.usecase.GetProductByKeyword(model.GetProductByKeywordRequest{Keyword: req.Keyword})
	var products []model.Product
	if resProduct.Status == http.StatusOK {
		products, _ = resProduct.Data["products"].([]model.Product)
	}

	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"Message":  message,
		"Products": products,
		"Amount":   len(products),
		"Search":   req.Keyword,
	})
}
