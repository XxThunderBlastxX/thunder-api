package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Link        string
	Description string
	Stacks      []TechStack `gorm:"foreignKey:id"`
}

type TechStack struct {
	gorm.Model
	Name string
}

type ProjectsRepository interface {
	AddProject(proj *Project) error
	GetProjects() ([]*Project, error)
	RemoveProject(id int) error
}

type ProjectsService interface {
	AddProject(proj *Project) error
	GetProjects() ([]*Project, error)
	RemoveProject(id int) error
}
