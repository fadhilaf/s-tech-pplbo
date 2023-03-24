package usecase

import (
	"net/http"
	"os"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
)

func (usecase *authUsecaseImpl) AdminLogin(req model.AdminLoginRequest) model.WebServiceResponse {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if req.Email != adminEmail {
		return utils.ToWebServiceResponse("Email salah", http.StatusUnauthorized, nil)
	}

	if req.Password != adminPassword {
		return utils.ToWebServiceResponse("Password salah", http.StatusUnauthorized, nil)
	}

	return utils.ToWebServiceResponse("Berhasil masuk sebagai admin", http.StatusOK, nil)
}
