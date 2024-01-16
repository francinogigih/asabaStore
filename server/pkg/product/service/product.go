package service

import (
	"asaba/pkg/product/core"
)

type ProductService struct {
	productRepository core.ProductRepository
}

func NewProductService(repo core.ProductRepository) ProductService {
	return ProductService{
		repo,
	}
}

func (s ProductService) Create(product *Product) (string, error) {
	data := core.Product{
		Code:        product.Code,
		Name:        product.Name,
		Total:       product.Total,
		Description: product.Description,
		Active:      product.Active,
	}
	id, err := s.productRepository.Create(&data)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s ProductService) Update(product *Product) error {
	data := core.Product{
		Name:        product.Name,
		Total:       product.Total,
		Description: product.Description,
		Active:      product.Active,
	}
	return s.productRepository.Update(&data, product.Code)
}
