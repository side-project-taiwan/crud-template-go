package usecase

import "spt/internal/service"

type ProjectUsecase interface {
	ListProjects() error
}

type projectUsecase struct {
	service service.ProjectService
}

func NewProjectUsecase(service service.ProjectService) ProjectUsecase {
	return &projectUsecase{service: service}
}

func (p *projectUsecase) ListProjects() error {
	return p.service.FetchProjects()
}
