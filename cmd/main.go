package main

import (
  "log"
  "fmt"

  appEnv "github.com/FadhilAF/s-tech-pplbo/common/env"
  mysqlCommon "github.com/FadhilAF/s-tech-pplbo/common/mysql"
  httpCommon "github.com/FadhilAF/s-tech-pplbo/common/http"

  landingPageDelivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/landing"
)

func main() {
  httpServer := httpCommon.NewHTTPServer() 

  config := appEnv.LoadConfig(".env");

  httpServer.Router.GET("/", landingPageDelivery.Render)

  db := mysqlCommon.Start(config.MysqlUrl);
  //sementara Close dulu biar compiler golang dk ngamuk
  db.Close();
  //apus line close pocok kalo nk lnjut koding

  log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%v", config.AppPort))) //kalo print, keluar di terminal, kalo Sprint return string
}
