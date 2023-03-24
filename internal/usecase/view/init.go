package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
)

type ViewUsecase interface {
	GetUserById(model.GetUserByIdRequest) model.WebServiceResponse
}

var _ ViewUsecase = &viewUsecaseImpl{}

func NewViewUsecase(store repository.Store) ViewUsecase {
	return &viewUsecaseImpl{
		Store: store,
	}
}

type viewUsecaseImpl struct {
	repository.Store
}
