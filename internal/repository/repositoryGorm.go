package repository

import (
	"sample/internal/repository/model"

	"gorm.io/gorm"
)

type RepositoryGorm struct {
	Database *gorm.DB
}

func NewRepositoryGorm(_database *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{
		Database: _database,
	}
}

func (r *RepositoryGorm) Signup(data *model.User) error {

	return nil
}
