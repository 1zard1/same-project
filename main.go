package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()
	s.Use(MiddleWare)
	s.GET("/status", Handler)
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	d := time.Date(2024, time.November, 30, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	s := fmt.Sprintf("Time until 30's November: %d", int64(dur.Hours())/24)
	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		value := c.Request().Header.Get("User-Role")
		if value == "admin" {
			log.Println("red button user detected")
		}
		err := next(c)
		if err != nil {
			return err
		}
		return nil
	}
}
