package model

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
)

type ProjectsResponse struct {
	Projects *[]projectDb.Project `json:"projects"`
}
