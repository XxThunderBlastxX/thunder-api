package models

type Project struct {
	Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Stacks      []string `json:"stacks"`
}

func NewProject(name, desc, link string, stacks []string) *Project {
	return &Project{
		Name:        name,
		Description: desc,
		Link:        link,
		Stacks:      stacks,
	}
}
