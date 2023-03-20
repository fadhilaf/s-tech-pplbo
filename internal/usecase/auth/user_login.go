package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *authUsecaseImpl) UserLogin(req model.UserLoginRequest) model.WebServiceResponse {
	user, err := usecase.Store.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return utils.ToWebServiceResponse("Email belum terdaftar", http.StatusNotFound, nil)
	}

	//belom sudah
	if err := utils.ComparePassword(req.Password, user.PasswordHash); err != nil {
		return utils.ToWebServiceResponse("Password salah", http.StatusUnauthorized, nil)
	}

	return utils.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
		"user": user,
	})
}
