package main

import (
	"github.com/FadhilAF/s-tech-pplbo/config"

	"github.com/FadhilAF/s-tech-pplbo/common/postgres"

	"github.com/FadhilAF/s-tech-pplbo/internal/app"
)

func main() {
	appConfig := config.LoadConfig(".env")

	postgresDb = postgres.Start("../common/postgres/migration", appConfig.PostgresUrl)

}
