package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type ViewUsecase interface {
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
