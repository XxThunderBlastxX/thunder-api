package project

type Repository interface {
	Create(Project) (Project, error)
	GetByID(string) (Project, error)
	GetAll() ([]Project, error)
	Update(string, Project) (Project, error)
	Delete(string) error
}
