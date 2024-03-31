package repository

import (
	"sample/internal/model"

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

func (r *RepositoryGorm) Signup(data *model.SignupRequest) error {

	return nil
}
