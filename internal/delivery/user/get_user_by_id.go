package delivery

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *userHandler) GetUserById(ctx *gin.Context) {
	var req model.GetUserRequest

	ok := utils.BindURIAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.GetUser(req)

	ctx.JSON(res.Status, res)
}
