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

	//convert productId dari string ke uuid
	productId, err := uuid.Parse(reqForm.ProductID)
	if err != nil {
		return
	}

	// Ambil data user
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		return
	}

	req := model.CreateOrderRequest{
		UserID:      userId,
		ProductID:   productId,
		Quantity:    reqForm.Quantity,
		Description: reqForm.Description,
	}
	res := handler.usecase.CreateOrder(req)
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC

	var location url.URL
	location = url.URL{Path: "/beli", RawQuery: "id=" + reqForm.ProductID}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/pesanan"}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
