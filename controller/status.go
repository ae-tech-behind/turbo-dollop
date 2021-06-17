package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

//go:generate mockgen -destination=./mocks/mock_status.go -package=mocks github.com/eiizu/go-service/controller StatusUseCase

// StatusUseCase -
type StatusUseCase interface {
	Statusz() (string, error)
	Healthz() (string, error)
}

// Status -
type Status struct {
	UseCase StatusUseCase
}

// NewStatus -
func NewStatus(uc StatusUseCase) *Status {
	return &Status{
		UseCase: uc,
	}
}

// HandlerStatusz -
func (c *Status) HandlerStatusz(eCtx echo.Context) error {
	resp, err := c.UseCase.Statusz()
	if err != nil {
		return eCtx.String(http.StatusBadRequest, "something went wrong")
	}

	return eCtx.JSON(http.StatusOK, resp)
}

// HandlerHealthz -
func (c *Status) HandlerHealthz(eCtx echo.Context) error {
	resp, err := c.UseCase.Healthz()
	if err != nil {
		return eCtx.String(http.StatusBadRequest, "something went wrong")
	}

	return eCtx.JSON(http.StatusOK, resp)
}
