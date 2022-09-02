package product

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return AllProducts
}

func (s *Service) Get(index int) (*Product, error) {
	if index > len(AllProducts)-1 {
		return nil, fmt.Errorf("out bound")
	}
	return &AllProducts[index], nil
}
