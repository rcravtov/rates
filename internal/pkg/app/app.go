package app

import (
	"log"
	"rates/internal/app/endpoint"
	"rates/internal/app/middleware"
	"rates/internal/app/repository"
	"rates/internal/app/service"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type App struct {
	endpoint  *endpoint.Endpoint
	service   *service.Service
	ginEngine *gin.Engine
	store     *repository.Store
}

func New() (*App, error) {

	app := App{}

	var err error
	app.store, err = repository.New()
	if err != nil {
		return &app, err
	}

	app.service = service.New(app.store)
	app.endpoint = endpoint.New(app.service)

	app.ginEngine = gin.Default()
	app.ginEngine.Use(middleware.CorsMiddleware())

	app.ginEngine.Use(static.ServeRoot("/", "/frontend"))
	app.ginEngine.GET("/status", app.endpoint.Status)
	app.ginEngine.GET("/currencies", app.endpoint.GetCurrencies)
	app.ginEngine.GET("/rates", app.endpoint.GetRates)
	//app.ginEngine.GET("/import", app.endpoint.ImportRates)

	return &app, nil

}

func (a *App) Run() error {

	log.Println("server running")
	err := a.ginEngine.Run()

	if err != nil {
		return err
	}

	return nil

}
