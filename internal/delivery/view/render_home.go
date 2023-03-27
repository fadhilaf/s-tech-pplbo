package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *viewHandler) RenderHome(c *gin.Context) {
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

	// Ambil data user
	userId := utils.GetUserIdFromContext(c)
	var name string
	if userId != uuid.Nil {

		resUser := handler.usecase.GetUserById(model.GetUserByIdRequest{ID: userId})
		if resUser.Status == http.StatusOK {

			user, ok := resUser.Data["user"].(model.User)
			if ok {

				name = user.Name
			}
		}
	}

	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"Message":  message,
		"Name":     name,
		"Products": products,
		"Amount":   len(products),
		"Search":   req.Keyword,
	})
}
