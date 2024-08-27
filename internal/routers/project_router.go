package routers

import (
	"github.com/gin-gonic/gin"
	"spt/internal/db"
	"spt/internal/handlers"
	"spt/internal/repository"
	"spt/internal/service"
	"spt/internal/usecase"
)

func RegisterProjectRoutes(r *gin.RouterGroup) {
	projectHandlers := setupDependencies()

	r.GET("/project", projectHandlers.GetProjectList)

}

func setupDependencies() handlers.ProjectHandler {
	db := db.Instance()
	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectUsecase := usecase.NewProjectUsecase(projectService)
	projectHandlers := handlers.NewProjectHandler(projectUsecase)
	return projectHandlers
}
