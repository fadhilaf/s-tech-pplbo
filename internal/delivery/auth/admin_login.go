package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *authHandler) AdminLogin(ctx *gin.Context) {
	var req model.AdminLoginRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.AdminLogin(req)
	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya HTML
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/admin"}

	if res.Status == http.StatusOK {
		utils.SaveAdminToSession(ctx, true)
		location = url.URL{Path: "/admin/dashboard"}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
