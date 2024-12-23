package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}
type Endpoint struct {
	svc Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		svc: svc,
	}
}

func (e *Endpoint) Status(c echo.Context) error {
	d := e.svc.DaysLeft()
	s := fmt.Sprintf("Time until 30's November: %d", d)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
