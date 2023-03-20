package delivery

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *authHandler) Logout(ctx *gin.Context) {
	utils.RemoveAuthSession(ctx)

	res := utils.ToWebServiceResponse("Logout success", http.StatusOK, nil)

	ctx.JSON(res.Status, res)
}
