package usecase

import (
	"strings"
)

type SomeService interface {
	GetParam() string
}

// Something -
type Something struct {
	SomeService SomeService
}

// NewSomething -
func NewSomething(s SomeService) *Something {
	return &Something{
		SomeService: s,
	}
}

// DoSomething -
func (uc *Something) DoSomething(data string) (map[string]int, error) {
	resp := map[string]int{}

	str := data + uc.SomeService.GetParam()
	arr := strings.Split(str, " ")

	for _, value := range arr {
		resp[value]++
	}

	return resp, nil
}
