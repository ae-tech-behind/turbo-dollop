package service

// SomeService -
type SomeService struct {
	Param string
}

// NewSomeService -
func NewSomeService(param string) *SomeService {
	return &SomeService{
		Param: param,
	}
}

// GetParam -
func (s *SomeService) GetParam() string {
	return s.Param
}
