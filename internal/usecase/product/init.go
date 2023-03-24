package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(req model.CreateProductRequest) model.WebServiceResponse
}

var _ ProductUsecase = &productUsecaseImpl{}

func NewProductUsecase(store repository.Store) ProductUsecase {
	return &productUsecaseImpl{
		Store: store,
	}
}

type productUsecaseImpl struct {
	repository.Store
}
