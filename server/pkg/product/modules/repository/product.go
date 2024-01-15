package repository

import (
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
