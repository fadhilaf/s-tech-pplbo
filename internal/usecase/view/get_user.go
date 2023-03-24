package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"

	utils "github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *viewUsecaseImpl) GetUserById(req model.GetUserByIdRequest) model.WebServiceResponse {
	userDb, err := usecase.Store.GetUserById(context.Background(), req.ID)
	if err != nil {
		return utils.ToWebServiceResponse("User tidak ditemukan", http.StatusNotFound, nil)
	}

	user := model.User{ID: userDb.ID, Name: userDb.Name, Email: userDb.Email, Address: userDb.Address, Phone: userDb.Phone}

	return utils.ToWebServiceResponse("Berhasil mendapatkan user", http.StatusOK, gin.H{
		"user": user,
	})
}
