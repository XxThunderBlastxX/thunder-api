package repository

import (
	"context"

	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
)

type projectsRepository struct {
	ctx   context.Context
	query *projectDb.Queries
}

func NewProjectsRepository(ctx context.Context, query *projectDb.Queries) domain.ProjectsRepository {
	return &projectsRepository{
		ctx:   ctx,
		query: query,
	}
}

func (p *projectsRepository) CreateProject(projParam *projectDb.CreateProjectParams) error {
	_, err := p.query.CreateProject(p.ctx, *projParam)
	if err != nil {
		return err
	}

	return nil
}

func (p *projectsRepository) ListProjects() (*[]projectDb.Project, error) {
	projList, err := p.query.ListProjects(p.ctx)
	if err != nil {
		return nil, err
	}

	return &projList, nil
}

func (p *projectsRepository) RemoveProjectById(id int32) error {
	err := p.query.DeleteProjectById(p.ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *projectsRepository) RemoveProjectByName(name string) error {
	err := p.query.DeleteProjectByName(p.ctx, name)
	if err != nil {
		return err
	}

	return nil
}
