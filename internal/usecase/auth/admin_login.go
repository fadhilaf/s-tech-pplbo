package usecase

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/config"
)

func (usecase *authUsecaseImpl) AdminLogin(req model.UserLoginRequest) model.WebServiceResponse {
	config := config.LoadConfig(".env")

	if req.Email != config.AdminEmail {
		return utils.ToWebServiceResponse("Email salah", http.StatusUnauthorized, nil)
	}

	if req.Password != config.AdminPassword {
		return utils.ToWebServiceResponse("Password salah", http.StatusUnauthorized, nil)
	}

	return utils.ToWebServiceResponse("Berhasil masuk sebagai admin", http.StatusOK, nil)
}
