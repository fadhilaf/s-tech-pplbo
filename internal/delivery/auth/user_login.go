package delivery

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *authHandler) UserLogin(ctx *gin.Context) {
	var req model.UserLoginRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.UserLogin(req)

	user, ok := res.Data.(model.User)
	if ok {
		utils.SaveUserToSession(ctx, user)
	}

	ctx.JSON(res.Status, res)
}
