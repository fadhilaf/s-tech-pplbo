package main

import (
	"fmt"
	"log"

	appEnv "github.com/FadhilAF/s-tech-pplbo/common/env"
	httpCommon "github.com/FadhilAF/s-tech-pplbo/common/http"
	postgresCommon "github.com/FadhilAF/s-tech-pplbo/common/postgres"

	landingPageDelivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/landing"
)

func main() {
	config := appEnv.LoadConfig(".env")

	httpServer := httpCommon.NewHTTPServer()
	db := postgresCommon.Start("file://config/db/migration", config.PostgresUrl)

	httpServer.Router.LoadHTMLGlob("internal/template/*.gohtml")

	httpServer.Router.GET("/", landingPageDelivery.Render)

	//sementara Close dulu biar compiler golang dk ngamuk
	db.Close()
	//apus line close pocok kalo nk lnjut koding

	log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%v", config.AppPort))) //kalo print, keluar di terminal, kalo Sprint return string
}
