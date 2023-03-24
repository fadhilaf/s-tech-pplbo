package usecase

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type UserUsecase interface {
	CreateUser(model.CreateUserRequest) model.WebServiceResponse
	// DeleteUser(model.DeleteUserRequest) model.WebServiceResponse
	// ListUser() model.WebServiceResponse
	// UpdateUser(model.UpdateUserRequest) model.WebServiceResponse
}

var _ UserUsecase = &userUsecaseImpl{}

func NewUserUsecase(store repository.Store) UserUsecase {
	return &userUsecaseImpl{
		Store: store,
	}
}

type userUsecaseImpl struct {
	repository.Store
}
