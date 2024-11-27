package place

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (h *Service) GetByID(id int) error {
	return errors.New("not implemented")
}
