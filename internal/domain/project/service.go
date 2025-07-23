package project

type Service interface {
	CreateProject(Project) (Project, error)
	GetAllProjects() ([]Project, error)
	GetProjectByID(string) (Project, error)
	UpdateProject(string, Project) (Project, error)
	DeleteProject(string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProject(project Project) (Project, error) {
	return s.repo.Create(project)
}

func (s *service) GetAllProjects() ([]Project, error) {
	return s.repo.GetAll()
}

func (s *service) GetProjectByID(id string) (Project, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateProject(id string, project Project) (Project, error) {
	return s.repo.Update(id, project)
}

func (s *service) DeleteProject(id string) error {
	return s.repo.Delete(id)
}
