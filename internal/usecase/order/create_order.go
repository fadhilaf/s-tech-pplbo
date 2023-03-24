package usecase

import (
	"context"

	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	repositoryModel "github.com/FadhilAF/s-tech-pplbo/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *orderUsecaseImpl) CreateOrder(req model.CreateOrderRequest) model.WebServiceResponse {
	product, err := usecase.Store.GetProductById(context.Background(), req.Form.ProductID)
	if err != nil {
		return utils.ToWebServiceResponse("Produk tidak ditemukan", http.StatusNotFound, nil)
	}

	if product.Stock < req.Form.Quantity {
		return utils.ToWebServiceResponse("Stok produk tidak mencukupi", http.StatusBadRequest, nil)
	}

	order, err := usecase.Store.CreateOrder(context.Background(), repositoryModel.CreateOrderParams{
		UserID:      req.UserID,
		ProductID:   req.Form.ProductID,
		Quantity:    req.Form.Quantity,
		Description: req.Form.Description,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan order ke database", http.StatusInternalServerError, nil)
	}

	updatedProduct, err := usecase.Store.UpdateProductStock(context.Background(), repositoryModel.UpdateProductStockParams{
		ID:    req.Form.ProductID,
		Stock: product.Stock - req.Form.Quantity,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengganti stok produk ke database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil memasukkan order ke database", http.StatusCreated, gin.H{
		"order":   order,
		"product": updatedProduct,
	})
}
