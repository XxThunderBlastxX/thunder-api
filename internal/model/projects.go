package model

import "github.com/XxThunderBlastxX/thunder-api/internal/domain"

type ProjectsResponse struct {
	Projects *[]domain.Project `json:"projects"`
}
