package app

type deliveries struct {
  user user_delivery.UserDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.user = division_delivery.NewDivisionDelivery(app.usecase.division)

	app.delivery = deliveries
}
