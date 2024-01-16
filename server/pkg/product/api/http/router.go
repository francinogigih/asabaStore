package http

import (
	"github.com/labstack/echo/v4"
)

// RegisterPath Register V1 API path
func RegisterPath(e *echo.Echo, h *Handler) {
	if h == nil {
		panic("item controller cannot be nil")
	}
	e.POST("v1/products", h.Create)
	e.PUT("v1/products/:code", h.Update)
}
