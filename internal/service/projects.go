package service

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
)

type projectsService struct {
	ProjectsRepo domain.ProjectsRepository
}

func NewProjectsService(projectsRepo domain.ProjectsRepository) domain.ProjectsService {
	return &projectsService{
		ProjectsRepo: projectsRepo,
	}
}

func (p *projectsService) CreateProject(project *projectDb.Project) error {
	projParam := &projectDb.CreateProjectParams{
		Name:        project.Name,
		Description: project.Description,
		Link:        project.Link,
		Stacks:      project.Stacks,
	}

	return p.ProjectsRepo.CreateProject(projParam)
}

func (p *projectsService) ListProjects() (*[]projectDb.Project, error) {
	return p.ProjectsRepo.ListProjects()

}

func (p *projectsService) RemoveProjectById(id int32) error {
	return p.ProjectsRepo.RemoveProjectById(id)
}

func (p *projectsService) RemoveProjectByName(name string) error {
	return p.ProjectsRepo.RemoveProjectByName(name)
}
