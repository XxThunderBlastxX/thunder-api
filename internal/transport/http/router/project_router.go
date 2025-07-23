package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/domain/project"
	"github.com/XxThunderBlastxX/thunder-api/internal/infrastructure/postgres"
	"github.com/XxThunderBlastxX/thunder-api/internal/transport/http/handler"
)

func (r *Router) projectRouter() {
	grp := r.app.Group("/projects")

	repo := postgres.NewProjectRepo(r.app.PgConn)
	s := project.NewService(repo)

	h := handler.NewProjectHandler(s)

	{
		grp.Post("", h.CreateProject())
		grp.Get("", h.GetAllProjects())
		grp.Get("/:id", h.GetProjectByID())
		grp.Put("/:id", h.UpdateProject())
		grp.Delete("/:id", h.DeleteProject())
	}
}
