package database

import (
	"fmt"
	"sample/configs"
	"sample/internal/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	util.PrintLogWithColor("Enter NewGormDB log")
	fmt.Println("configs:", configs.DB_TYPE)

	newGormDB, err := gorm.Open(mysql.Open(configs.DB_URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := newGormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return newGormDB, nil
}
