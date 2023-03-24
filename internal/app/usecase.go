package app

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"

	auth_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/auth"
	product_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/product"
	user_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/user"
	view_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/view"
)

type usecases struct {
	auth    auth_usecase.AuthUsecase
	user    user_usecase.UserUsecase
	product product_usecase.ProductUsecase
	view    view_usecase.ViewUsecase
}

func (app *App) initUsecase() {
	store := repository.NewPostgresStore(app.db)

	var usecases usecases
	usecases.auth = auth_usecase.NewAuthUsecase(store)
	usecases.user = user_usecase.NewUserUsecase(store)
	usecases.product = product_usecase.NewProductUsecase(store)
	usecases.view = view_usecase.NewViewUsecase(store)

	app.usecase = usecases
}
