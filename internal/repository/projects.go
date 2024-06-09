package repository

import (
	"fmt"

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

func (p *projectsRepository) ListProjects() (*[]domain.Project, error) {
	projects := &[]domain.Project{}
	trx := p.db.Preload("Stacks").Find(&projects)
	if trx.Error != nil {
		return nil, trx.Error
	}

	return projects, nil
}

func (p *projectsRepository) RemoveProjectById(id string) error {
	trx := p.db.Where("id = ?", id).Delete(&domain.Project{})
	if trx.Error != nil {
		return trx.Error
	}

	if trx.RowsAffected == 0 {
		return fmt.Errorf("no project found with id %s", id)
	}

	return nil
}

func (p *projectsRepository) RemoveProjectByName(name string) error {
	trx := p.db.Where("name = ?", name).Delete(&domain.Project{})
	if trx.Error != nil {
		return trx.Error
	}

	if trx.RowsAffected == 0 {
		return fmt.Errorf("no project found with name %s", name)
	}

	return nil
}
