package http

import (
	"asaba/pkg/product/api/http/request"
	"asaba/pkg/product/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.ProductService
}

// NewHandler Construct user API handler
func NewHandler(service service.ProductService) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) Create(c echo.Context) error {
	req := new(request.CreateProduct)
	c.Bind(req)

	dto := newProductData(*req)
	code, err := h.service.Create(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "resource error.",
			"message": "could not insert data.",
		})
	}

	return c.JSON(http.StatusOK, code)
}
