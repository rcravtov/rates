package app

import (
	"log"
	"rates/internal/app/endpoint"
	"rates/internal/app/middleware"
	"rates/internal/app/repository"
	"rates/internal/app/service"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

	app.service = service.New(app.store, cron.New())
	app.endpoint = endpoint.New(app.service)

	app.ginEngine = gin.Default()
	app.ginEngine.Use(middleware.CorsMiddleware())

	app.ginEngine.Use(static.ServeRoot("/", "/frontend"))
	app.ginEngine.GET("/api/status", app.endpoint.Status)
	app.ginEngine.GET("/api/currencies", app.endpoint.GetCurrencies)
	app.ginEngine.GET("/api/rates", app.endpoint.GetRates)
	app.ginEngine.POST("/api/auth", app.endpoint.Auth)

	protected := app.ginEngine.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/auth_settings", app.endpoint.GetAuthSettings)
	protected.POST("/auth_settings", app.endpoint.SetAuthSettings)
	protected.GET("/settings", app.endpoint.GetSettings)
	protected.POST("/settings", app.endpoint.SetSettings)
	protected.GET("/import", app.endpoint.ImportRates)
	protected.GET("/import_logs", app.endpoint.GetImportLogs)

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
