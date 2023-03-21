package delivery

import (
	"fmt"
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

	//Gaya MVC
	fmt.Println(res)

	utils.SaveInputData(ctx, model.InputData{Message: res.Message})

	location := url.URL{Path: "/register"}
	ctx.Redirect(res.Status, location.RequestURI())
}
