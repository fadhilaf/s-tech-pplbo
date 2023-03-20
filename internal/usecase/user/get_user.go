package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"

	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *userUsecaseImpl) GetUser(req model.GetUserRequest) model.WebServiceResponse {
	user, err := usecase.Store.GetUserById(context.Background(), req.ID)
	if err != nil {
		return utils.ToWebServiceResponse("Error ketika mencari user di database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan user", http.StatusOK, gin.H{
		"user": user,
	})
}
