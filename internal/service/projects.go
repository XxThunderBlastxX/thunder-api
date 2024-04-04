package service

import "github.com/XxThunderBlastxX/thunder-api/internal/domain"

type projectsService struct {
	ProjectsRepo domain.ProjectsRepository
}

func NewProjectsService(projectsRepo domain.ProjectsRepository) domain.ProjectsService {
	return &projectsService{
		ProjectsRepo: projectsRepo,
	}
}

func (p *projectsService) AddProject(proj *domain.Project) error {
	return p.ProjectsRepo.AddProject(proj)
}

func (p *projectsService) GetProjects() (*[]domain.Project, error) {
	return p.ProjectsRepo.GetProjects()

}

func (p *projectsService) RemoveProjectById(id string) error {
	return p.ProjectsRepo.RemoveProjectById(id)
}

func (p *projectsService) RemoveProjectByName(name string) error {
	return p.ProjectsRepo.RemoveProjectByName(name)
}
