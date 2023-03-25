package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/google/uuid"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) DeliveredOrder(ctx *gin.Context) {
	var reqForm model.UpdateOrderStatusFormRequest

	ok := utils.BindFormAndValidate(ctx, &reqForm)
	if !ok {
		return
	}

	//convert orderId dari string ke uuid
	orderId, err := uuid.Parse(reqForm.ID)
	if err != nil {
		return
	}

	// Ambil data user
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		return
	}

	req := model.UpdateOrderStatusDeliveredRequest{
		ID:     orderId,
		UserID: userId,
	}
	res := handler.usecase.DeliveredOrder(req)
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC

	location := url.URL{Path: "/pesanan"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
