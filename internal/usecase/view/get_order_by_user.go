package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"

	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *viewUsecaseImpl) GetOrderByUserId(req model.GetOrderByUserIdRequest) model.WebServiceResponse {
	ordersDb, err := usecase.Store.GetOrdersByUserId(context.Background(), req.UserID)
	if err != nil {
		return utils.ToWebServiceResponse("Orders tidak ditemukan", http.StatusNotFound, nil)
	}

	orders := make([]model.Order, len(ordersDb))

	for i, order := range ordersDb {
		product, err := usecase.Store.GetProductById(context.Background(), order.ProductID)
		if err != nil {
			return utils.ToWebServiceResponse("Product tidak ditemukan di database", http.StatusNotFound, nil)
		}

		user, err := usecase.Store.GetUserById(context.Background(), order.UserID)
		if err != nil {
			return utils.ToWebServiceResponse("User tidak ditemukan di database", http.StatusNotFound, nil)
		}

		orders[i] = model.Order{
			ID:           order.ID,
			ProductID:    product.ID,
			ProductName:  product.Name,
			BuyerID:      user.ID,
			BuyerName:    user.Name,
			BuyerAddress: user.Address,
			BuyerPhone:   user.Phone,
			Quantity:     order.Quantity,
			Status:       order.Status.(string),
			Description:  order.Description,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan orders", http.StatusOK, gin.H{
		"orders": orders,
	})
}
