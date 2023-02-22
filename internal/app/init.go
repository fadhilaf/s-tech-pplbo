package app

import (
	"database/sql"
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/config"

	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type App struct {
	Config   config.Config
	delivery deliveries
	usecase  usecases
	store    repository.Store
}

func (app *App) createHandlers() *gin.Engine {
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Accept")

	corsCfg.AllowAllOrigins = true

	router.Use(cors.New(corsCfg))

	viewRouterGroup := router.Group("/")
	apiRouterGroup := router.Group("/api")

	app.viewHandler(viewRouterGroup)
	app.apiHandler(apiRouterGroup)

	// mesin untuk menampilkan semua routes yang ada?
	routes := router.Routes()
	if gin.Mode() == gin.DebugMode {
		fmt.Println()
		for _, v := range routes {
			path := color.InBold(v.Path)
			method := color.InYellow(fmt.Sprintf("%-6s", v.Method))
			fmt.Println(method, path)
		}
		fmt.Println()
	}
	// mesin untuk menampilkan semua routes yang ada? selesai

	return router
}

func (app *App) StartServer() {
	if app.Config.Env == config.EnvProd {
		fmt.Println(
			color.Ize(color.Yellow, color.InBold("\nAPP RUN IN PRODUCTION MODE\n")),
		)
	} else {
		fmt.Println(
			color.Ize(color.Red, color.InBold("\nAPP RUN IN DEVELOPMENT MODE\n")),
		)
	}

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validations.InitValidations(validator)
	}
	router := app.createHandlers()
	address := fmt.Sprintf("%s:%s", app.Config.AppHost, app.Config.AppPort)
	log.Printf("Server listening on %v\n", address)

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Cannot start server %v\n", err)
		}
	}()

	<-osSignalChan
	err := srv.Close()
	if err != nil {
		log.Fatalf("cannot shutdown server %v", err)
	}
	fmt.Println()
	log.Println("Server exiting")
}

func New(config config.Config, db *sql.DB) App {
	app := App{}

	app.Config = config

	app.store = repository.NewPostgresStore(db)
	app.initUsecase()
	app.initDelivery()

	return app
}
