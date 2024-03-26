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

func (p *projectsService) RemoveProject(id int) error {
	return p.ProjectsRepo.RemoveProject(id)
}
