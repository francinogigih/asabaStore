package http

import (
	common "asaba/pkg/common/http"
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
	req := new(request.Product)
	c.Bind(req)

	dto := newProductData(*req)
	code, err := h.service.Create(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "resource error.",
			"message": "could not insert data.",
		})
	}

	resp := common.NewCreatedSuccessResponse(code)
	return c.JSON(resp.Code, resp)
}

func (h *Handler) Update(c echo.Context) error {
	req := new(request.Product)
	req.Code = c.Param("code")
	c.Bind(req)

	dto := newUpdateProductData(*req)
	if err := h.service.Update(dto); err != nil {
		errResp := common.RenderErrorResponse(err)
		return c.JSON(errResp.Code, errResp)
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponse())
}
