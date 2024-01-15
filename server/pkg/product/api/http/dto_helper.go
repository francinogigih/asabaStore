package http

import (
	"asaba/pkg/product/api/http/request"
	"asaba/pkg/product/service"
)

func newProductData(params request.CreateProduct) *service.Product {
	return &service.Product{
		Code:        params.Code,
		Name:        params.Name,
		Total:       params.Total,
		Description: params.Description,
		Active:      params.Active,
	}
}
