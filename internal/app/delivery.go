package app

import (
	user_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/user"
	view_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/view"
)

type deliveries struct {
	user user_delivery.UserDelivery
	view view_delivery.ViewDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.user = user_delivery.NewUserDelivery(app.usecase.user)
	deliveries.view = view_delivery.NewViewDelivery(app.usecase.view)

	app.delivery = deliveries
}
