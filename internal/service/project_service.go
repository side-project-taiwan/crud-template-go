package service

import "spt/internal/repository"

type ProjectService interface {
	FetchProjects() error
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{
		repo: repo,
	}
}

func (p *projectService) FetchProjects() error {
	return p.repo.FindAll()
}
