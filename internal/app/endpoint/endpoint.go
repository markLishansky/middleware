package endpoint

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()
	s := fmt.Sprintf("Days reamin: %d", d)


	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	
	return nil
}