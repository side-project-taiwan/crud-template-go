package repository

import (
	"errors"
	"spt/internal/db"
)

type ProjectRepository interface {
	FindAll() error
}

type projectRepository struct {
	dbService db.Service
}

func NewProjectRepository(dbService db.Service) ProjectRepository {
	return &projectRepository{
		dbService: dbService,
	}
}

func (p *projectRepository) FindAll() error {
	//TODO implement me
	return errors.New("not implemented")
}
