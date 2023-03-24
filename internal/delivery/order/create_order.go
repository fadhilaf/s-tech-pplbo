package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/google/uuid"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) CreateOrder(ctx *gin.Context) {
	var reqForm model.CreateOrderFormRequest

	ok := utils.BindFormAndValidate(ctx, &reqForm)
	if !ok {
		return
	}

	// Ambil data user
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		return
	}

	req := model.CreateOrderRequest{
		UserID: userId,
		Form:   reqForm,
	}

	res := handler.usecase.CreateOrder(req)
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC

	var location url.URL
	location = url.URL{Path: "/buy", RawQuery: "id=" + reqForm.ProductID.String()}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/order"}
	}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
