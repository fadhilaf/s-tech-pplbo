package app

import (
	auth_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/auth"
	product_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/product"
	user_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/user"
	view_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/view"
)

type deliveries struct {
	auth    auth_delivery.AuthDelivery
	user    user_delivery.UserDelivery
	product product_delivery.ProductDelivery
	view    view_delivery.ViewDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.auth = auth_delivery.NewAuthDelivery(app.usecase.auth)
	deliveries.user = user_delivery.NewUserDelivery(app.usecase.user)
	deliveries.product = product_delivery.NewProductDelivery(app.usecase.product)
	deliveries.view = view_delivery.NewViewDelivery(app.usecase.view)

	app.delivery = deliveries
}
