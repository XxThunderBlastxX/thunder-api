package postgres

import "github.com/XxThunderBlastxX/thunder-api/internal/domain/project"

type projectRepo struct {
	dbConn *ConnectionManager
}

func NewProjectRepo(dbConn *ConnectionManager) project.Repository {
	return &projectRepo{
		dbConn: dbConn,
	}
}

// Create implements project.Repository.
func (p *projectRepo) Create(project project.Project) (project.Project, error) {
	panic("unimplemented")
}

// Delete implements project.Repository.
func (p *projectRepo) Delete(id string) error {
	panic("unimplemented")
}

// GetAll implements project.Repository.
func (p *projectRepo) GetAll() ([]project.Project, error) {
	panic("unimplemented")
}

// GetByID implements project.Repository.
func (p *projectRepo) GetByID(id string) (project.Project, error) {
	panic("unimplemented")
}

// Update implements project.Repository.
func (p *projectRepo) Update(id string, project project.Project) (project.Project, error) {
	panic("unimplemented")
}
