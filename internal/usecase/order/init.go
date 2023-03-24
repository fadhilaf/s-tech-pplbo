package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type OrderUsecase interface {
	CreateOrder(model.CreateOrderRequest) model.WebServiceResponse
}

var _ OrderUsecase = &orderUsecaseImpl{}

func NewOrderUsecase(store repository.Store) OrderUsecase {
	return &orderUsecaseImpl{
		Store: store,
	}
}

type orderUsecaseImpl struct {
	repository.Store
}
