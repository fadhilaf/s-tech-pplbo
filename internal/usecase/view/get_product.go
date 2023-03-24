package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"

	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *viewUsecaseImpl) GetProduct() model.WebServiceResponse {
	productsDb, err := usecase.Store.GetProduct(context.Background())
	if err != nil {
		return utils.ToWebServiceResponse("User tidak ditemukan", http.StatusNotFound, nil)
	}

	products := make([]model.Product, len(productsDb))

	for i, product := range productsDb {
		products[i] = model.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			IsService:   product.IsService,
			Description: product.Description,
			Image:       product.Image,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan user", http.StatusOK, gin.H{
		"products": products,
	})
}
