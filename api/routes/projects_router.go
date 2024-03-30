package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func ProjectsRouter(router fiber.Router, db *gorm.DB) {
	projectsRepo := repository.NewProjectsRepository(db)
	projectsService := service.NewProjectsService(projectsRepo)
	ctr := controller.ProjectsController{ProjectsService: projectsService}

	proj := router.Group("/projects")

	proj.Post("/add", ctr.AddProject())
	proj.Get("/list", ctr.GetProjects())
	proj.Delete("/remove/:name", ctr.RemoveProject())
}
