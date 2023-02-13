package main

import (
	"log"

	httpCommon "github.com/FadhilAF/s-tech-pplbo/common/http"

	landingPageDelivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/landing"
)

func main() {

	httpServer := httpCommon.NewHTTPServer()

	httpServer.Router.LoadHTMLGlob("internal/template/*.gohtml")

	httpServer.Router.GET("/", landingPageDelivery.Render)

	//sementara Close dulu biar compiler golang dk ngamuk
	//apus line close pocok kalo nk lnjut koding

	log.Fatalln(httpServer.Router.Run()) //kalo print, keluar di terminal, kalo Sprint return string
}
