package usecase

import (
	"context"

	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	repositoryModel "github.com/FadhilAF/s-tech-pplbo/internal/repository/postgres/sqlc"

	passwordUtils "github.com/FadhilAF/s-tech-pplbo/common/password"
	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *userUsecaseImpl) CreateUser(req model.CreateUserRequest) model.WebServiceResponse {
	passwordHashManager := passwordUtils.NewPasswordHashManagerBcrypt()

	passwordHash, err := passwordHashManager.HashPassword(req.Password)
	if err != nil {
		return utils.ToWebServiceResponse("Fungsi hash password gagal", http.StatusInternalServerError, nil)
	}

	_, err = usecase.Store.GetUserByEmail(context.Background(), req.Email)
	if err == nil {
		return utils.ToWebServiceResponse("Email sudah terdaftar", http.StatusConflict, nil)
	}

	_, err = usecase.Store.CreateUser(context.Background(), repositoryModel.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		Address:      req.Address,
		Phone:        req.Phone,
	})

	if err != nil {
		return utils.ToWebServiceResponse("Menambah user ke db gagal", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Divisi berhasil dibuat", http.StatusCreated, nil)
}
