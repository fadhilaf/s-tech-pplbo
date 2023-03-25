package delivery

import (
	usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/view"

	"github.com/gin-gonic/gin"
)

type ViewDelivery interface {
	RenderHome(ctx *gin.Context)
	RenderRegister(ctx *gin.Context)
	RenderLogin(ctx *gin.Context)
	RenderAdmin(ctx *gin.Context)
	RenderDashboard(ctx *gin.Context)
	RenderPesan(ctx *gin.Context)
	RenderPesanan(ctx *gin.Context)
	RenderAdminPesanan(ctx *gin.Context)
	RenderTambah(ctx *gin.Context)
}

var _ ViewDelivery = &viewHandler{}

func NewViewDelivery(usecase usecase.ViewUsecase) ViewDelivery {
	return &viewHandler{
		usecase: usecase,
	}
}

type viewHandler struct {
	usecase usecase.ViewUsecase
}
