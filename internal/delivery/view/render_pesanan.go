package delivery

import (
	"log"
	"net/http"
	"os"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func (handler *viewHandler) RenderPesanan(c *gin.Context) {
	// Ambil data user
	userId := utils.GetUserIdFromContext(c)
	if userId == uuid.Nil {
		return
	}

	resUser := handler.usecase.GetUserById(model.GetUserByIdRequest{ID: userId})
	if resUser.Status != http.StatusOK {
		return
	}
	user, ok := resUser.Data["user"].(model.User)

	var name string
	if ok {
		name = user.Name
	}

	// Ambil data pesanan
	resOrder := handler.usecase.GetOrderByUserId(model.GetOrderByUserIdRequest{UserID: userId})
	var orders []model.Order
	if resOrder.Status == http.StatusOK {
		orders, _ = resOrder.Data["orders"].([]model.Order)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error ketika membaca .env: %s", err)
	}

	adminPhone := os.Getenv("ADMIN_PHONE")

	c.HTML(http.StatusOK, "pesanan.gohtml", gin.H{
		"Name":       name,
		"AdminPhone": adminPhone,
		"Orders":     orders,
	})
}
