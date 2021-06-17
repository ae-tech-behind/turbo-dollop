package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

//go:generate mockgen -destination=./mocks/mock_something.go -package=mocks github.com/eiizu/go-service/controller SomethingUseCase

// SomethingUseCase -
type SomethingUseCase interface {
	DoSomething(string) (map[string]int, error)
}

// Something -
type Something struct {
	UseCase SomethingUseCase
}

// Request -
type Request struct {
	Info string `json:"info"`
}

// NewSomething -
func NewSomething(uc SomethingUseCase) *Something {
	return &Something{
		UseCase: uc,
	}
}

// HandlerSomething -
func (c *Something) HandlerSomething(eCtx echo.Context) error {
	decoder := json.NewDecoder(eCtx.Request().Body)

	var data Request
	if err := decoder.Decode(&data); err != nil {
		return eCtx.String(http.StatusBadRequest, "invalid json")
	}

	if data.Info == "" {
		return eCtx.String(http.StatusBadRequest, "invalid info")
	}

	resp, err := c.UseCase.DoSomething(data.Info)
	if err != nil {
		return eCtx.String(http.StatusBadRequest, "something went wrong")
	}

	return eCtx.JSON(http.StatusOK, resp)
}
