package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/google/uuid"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) ProcessingOrder(ctx *gin.Context) {
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

	req := model.UpdateOrderStatusProcessingRequest{
		ID: orderId,
	}
	res := handler.usecase.ProcessingOrder(req)
	utils.SaveResponse(ctx, res.Message)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	location := url.URL{Path: "/admin/pesanan"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
