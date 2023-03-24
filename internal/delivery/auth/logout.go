package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/gin-gonic/gin"
)

func (handler *authHandler) Logout(ctx *gin.Context) {
	utils.RemoveAuthSession(ctx)

	res := utils.ToWebServiceResponse("Logout success", http.StatusOK, nil)

	//Gaya REST API
	// ctx.JSON(res.Status, res)

	//Gaya HTML
	utils.SaveResponse(ctx, res.Message)

	location := url.URL{Path: "/login"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
