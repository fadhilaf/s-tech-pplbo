package delivery

import (
	"net/http"
	"net/url"

	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *productHandler) CreateProduct(ctx *gin.Context) {
	var req model.CreateProductNoFileRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	// Simpan upload file ke folder assets/images
	filename, ok := utils.SaveFileFromForm(ctx, "image", "internal/template/assets/img/")
	if !ok {
		return
	}

	res := handler.usecase.CreateProduct(model.CreateProductRequest{
		NotFile: req,
		Image:   filename,
	})

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/admin/tambah"}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/admin/dashboard"}
	}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
