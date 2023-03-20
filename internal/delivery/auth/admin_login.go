package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *authHandler) AdminLogin(ctx *gin.Context) {
	var req model.UserLoginRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.UserLogin(req)

	if res.Status == http.StatusOK {
		utils.SaveAdminToSession(ctx, true)
	}

	ctx.JSON(res.Status, res)
}
