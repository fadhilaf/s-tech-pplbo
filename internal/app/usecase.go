package app

import (
	user_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/user"
	view_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/view"
)

type usecases struct {
	user user_usecase.UserUsecase
	view view_usecase.ViewUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
	usecases.user = user_usecase.NewUserUsecase(app.store)
	usecases.view = view_usecase.NewViewUsecase(app.store)

	app.usecase = usecases
}
