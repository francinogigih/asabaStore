package response

import "asaba/pkg/product/core"

type ProductResponse struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Total       int64  `json:"total"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type GetProductsResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Payload []ProductResponse `json:"payload"`
}

func NewGetProductsResponse(result []core.Product) *GetProductsResponse {
	var ResultResponse GetProductsResponse

	products := make([]ProductResponse, len(result))
	for i, product := range result {
		products[i] = ProductResponse{
			Code:        product.Code,
			Name:        product.Name,
			Total:       product.Total,
			Description: product.Description,
			Active:      product.Active,
		}
	}

	ResultResponse.Code = 200
	ResultResponse.Message = "Success"
	ResultResponse.Payload = products

	return &ResultResponse
}
