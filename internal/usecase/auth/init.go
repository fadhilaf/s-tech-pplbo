package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type AuthUsecase interface {
	UserLogin(model.LoginRequest) model.WebServiceResponse
	AdminLogin(model.LoginRequest) model.WebServiceResponse
}

var _ AuthUsecase = &authUsecaseImpl{}

func NewAuthUsecase(store repository.Store) AuthUsecase {
	return &authUsecaseImpl{
		Store: store,
	}
}

type authUsecaseImpl struct {
	repository.Store
}
