package service

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/database"
	"github.com/XxThunderBlastxX/thunder-api/internal/models"
)

type ProjectService struct {
	db *database.ConnectionManager
}

func NewProjectsService(db *database.ConnectionManager) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (p *ProjectService) CreateProject(project *models.Project) error {
	if err := p.db.Create(project).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) GetProjectByID(id string) (*models.Project, error) {
	var project models.Project
	if err := p.db.First(&project, id).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *ProjectService) UpdateProject(id string, project *models.Project) error {
	if err := p.db.Where("id = ?", id).Updates(project).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) DeleteProject(id string) error {
	if err := p.db.Delete(&models.Project{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) ListProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := p.db.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}
