package app

import (
	"fmt"
	"log"

	mw "github.com/1zard1/same-project/internal/app/MW"
	"github.com/1zard1/same-project/internal/app/endpoint"
	"github.com/1zard1/same-project/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
