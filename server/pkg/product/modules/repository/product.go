package repository

import (
	appErr "asaba/pkg/common/http"
	"asaba/pkg/product/core"
	"errors"

	"gorm.io/gorm"
)

type PostgresDBRepository struct {
	collection *gorm.DB
}

func NewPostgresDBRepository(db *gorm.DB) *PostgresDBRepository {
	return &PostgresDBRepository{
		db,
	}
}

type Product struct {
	Code        string `gorm:"index:code,unique"`
	Name        string `gorm:"index:name,unique"`
	Total       int64
	Description string
	Active      bool
}

func (repo *PostgresDBRepository) newProductData(product *core.Product) Product {
	return Product{
		Code:        product.Code,
		Name:        product.Name,
		Total:       product.Total,
		Description: product.Description,
		Active:      product.Active,
	}
}

func (repo *PostgresDBRepository) Create(data *core.Product) (string, error) {
	product := repo.newProductData(data)
	result := repo.collection.Create(&product)

	if result.Error != nil {
		return "", errors.New(appErr.ErrInternalServer)
	}
	return product.Code, nil
}

func (repo *PostgresDBRepository) Update(data *core.Product, code string) error {
	product := repo.newProductData(data)
	result := repo.collection.Where("code = ?", code).Updates(&product)

	if result.RowsAffected == 0 {
		return errors.New(appErr.ErrNotFound)
	}
	if result.Error != nil {
		return errors.New(appErr.ErrInternalServer)
	}
	return nil
}
