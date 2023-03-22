package main

import (
	"github.com/FadhilAF/s-tech-pplbo/common/env"

	"github.com/FadhilAF/s-tech-pplbo/common/postgres"

	"github.com/FadhilAF/s-tech-pplbo/internal/app"
)

func main() {
	appConfig := env.LoadConfig(".env")

	postgresDb := postgres.Start(appConfig.PostgresUrl)

	app := app.New(appConfig, postgresDb)

	app.StartServer()
}
