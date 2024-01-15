package main

import (
	"context"
	"fmt"

	"asaba/config"
	pg "asaba/pkg/common/modules/db"
	productAPIHttp "asaba/pkg/product/api/http"
	productRepository "asaba/pkg/product/modules/repository"
	productService "asaba/pkg/product/service"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/labstack/gommon/log"
)

func newProductService(db *gorm.DB) productService.ProductService {
	productRepository := productRepository.NewPostgresDBRepository(db)
	productService := productService.NewProductService(productRepository)
	return productService
}

func main() {
	db := pg.DatabaseConnection()

	productService := newProductService(db)
	productHandler := productAPIHttp.NewHandler(productService)

	ctx := context.Background()
	// define echo backend
	e := echo.New()
	e.Use(mw.Recover())

	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	productAPIHttp.RegisterPath(e, productHandler)

	address := fmt.Sprintf(":%d", config.GetConfig().AppPort)
	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
