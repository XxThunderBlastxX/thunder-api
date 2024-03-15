package repository

import (
	"gorm.io/gorm"

	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
)

type projectsRepository struct {
	db *gorm.DB
}

func NewProjectsRepository(db *gorm.DB) domain.ProjectsRepository {
	return &projectsRepository{
		db: db,
	}
}

func (p *projectsRepository) AddProject(proj *domain.Project) error {
	trx := p.db.Create(&proj)
	if trx.Error != nil {
		return trx.Error
	}

	return nil
}

func (p *projectsRepository) GetProjects() ([]*domain.Project, error) {
	var projects []*domain.Project
	trx := p.db.Find(&projects)
	if trx.Error != nil {
		return nil, trx.Error
	}

	return projects, nil
}

func (p *projectsRepository) RemoveProject(id int) error {
	trx := p.db.Delete(&domain.Project{}, id)
	if trx.Error != nil {
		return trx.Error
	}

	return nil
}
