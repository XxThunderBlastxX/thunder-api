package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/handler"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func (r *Router) projectRouter() {
	grp := r.app.Group("/projects")

	s := service.NewProjectsService(r.app.PgConn)
	h := handler.NewProjectHandler(s)

	{
		grp.Post("", h.CreateProject())
		grp.Get("", h.GetAllProjects())
		grp.Get("/:id", h.GetProjectByID())
		grp.Put("/:id", h.UpdateProject())
		grp.Delete("/:id", h.DeleteProject())
	}
}
