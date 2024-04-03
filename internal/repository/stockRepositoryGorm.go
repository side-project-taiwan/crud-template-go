package repository

import (
	"gorm.io/gorm"
)

type StockRepositoryGorm struct {
	Database *gorm.DB
}

func NewRepositoryGorm(_database *gorm.DB) *StockRepositoryGorm {
	return &StockRepositoryGorm{
		Database: _database,
	}
}
