package routes

import (
	"context"
	"database/sql"
	"github.com/XxThunderBlastxX/thunder-api/api/middleware"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func ProjectsRouter(router fiber.Router, db *sql.DB) {
	ctx := context.Background()
	dbQuery := projectDb.New(db)
	jwtMiddleware := middleware.NewJWTMiddleware()

	projectsRepo := repository.NewProjectsRepository(ctx, dbQuery)
	projectsService := service.NewProjectsService(projectsRepo)
	ctr := controller.ProjectsController{ProjectsService: projectsService}

	proj := router.Group("/projects")

	proj.Post("/add", jwtMiddleware, ctr.AddProject())
	proj.Delete("/remove", jwtMiddleware, ctr.RemoveProject())
	proj.Get("/list", ctr.ListProjects())
}
