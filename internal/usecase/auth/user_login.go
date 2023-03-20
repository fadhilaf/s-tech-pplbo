package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
)

func (usecase *authUsecaseImpl) UserLogin(req model.UserLoginRequest) model.WebServiceResponse {
	user, err := usecase.Store.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return model.WebServiceResponse{http.NotFound, "Email belum terdaftar", nil}
	}

	//belom sudah
}
