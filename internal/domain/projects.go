package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string      `json:"name"`
	Link        string      `json:"link"`
	Description string      `json:"description"`
	Stacks      []TechStack `gorm:"foreignKey:ProjectID" json:"stacks"`
}

type TechStack struct {
	gorm.Model
	Name      string `json:"name"`
	ProjectID uint   `gorm:"index"`
}

type ProjectsRepository interface {
	AddProject(proj *Project) error
	GetProjects() (*[]Project, error)
	RemoveProjectById(id string) error
	RemoveProjectByName(name string) error
}

type ProjectsService interface {
	AddProject(proj *Project) error
	GetProjects() (*[]Project, error)
	RemoveProjectById(id string) error
	RemoveProjectByName(name string) error
}
