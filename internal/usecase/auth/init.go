package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type AuthUsecase interface {
	UserLogin(model.UserLoginRequest) model.WebServiceResponse
	AdminLogin(model.UserLoginRequest) model.WebServiceResponse
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
