package domain

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
)

type ProjectsRepository interface {
	CreateProject(projParam *projectDb.CreateProjectParams) error
	ListProjects() (*[]projectDb.Project, error)
	RemoveProjectById(id int32) error
	RemoveProjectByName(name string) error
}

type ProjectsService interface {
	CreateProject(project *projectDb.Project) error
	ListProjects() (*[]projectDb.Project, error)
	RemoveProjectById(id int32) error
	RemoveProjectByName(name string) error
}
