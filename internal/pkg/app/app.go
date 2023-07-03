package app

import (
	"grisko-test-app/config"
	"grisko-test-app/internal/app/database"
	"grisko-test-app/internal/app/endpoint"
	"grisko-test-app/internal/app/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	db   *database.DataBase
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	DataBase, err := database.New("grisko", "products")
	if err != nil {
		panic(err)
	}

	a.db = DataBase

	a.e = endpoint.New(a.s, a.db)

	a.echo = echo.New()

	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	a.echo.GET("/products", a.e.ReadProduct)
	a.echo.POST("/products", a.e.CreateProduct)
	a.echo.DELETE("/products", a.e.DeleteProduct)
	a.echo.PUT("/products", a.e.UpdateProduct)

	return a, nil

}

func (a *App) Run() error {

	err := a.echo.Start(config.ServerPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server running...")
	return nil
}
