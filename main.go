package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()
	s.GET("/status", Handler)
	err := s.Start(":8080")]
	if err != nil {
		log.Fatal(err)
	}
}


func Handler(c echo.Context) error {
	err := c.String(http.StatusOK, "test")
	if err != nil {
		return err
	}
	return nil
}