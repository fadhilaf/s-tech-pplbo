package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *userHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}
	res := handler.usecase.CreateUser(req)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/login"}
	} else {
		location = url.URL{Path: "/register"}
	}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
