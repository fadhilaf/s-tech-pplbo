package app

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/repository"

	user_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/user"
	view_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/view"
)

type usecases struct {
	user user_usecase.UserUsecase
	view view_usecase.ViewUsecase
}

func (app *App) initUsecase() {
	store := repository.NewPostgresStore(app.db)

	var usecases usecases
	usecases.user = user_usecase.NewUserUsecase(store)
	usecases.view = view_usecase.NewViewUsecase(store)

	app.usecase = usecases
}
