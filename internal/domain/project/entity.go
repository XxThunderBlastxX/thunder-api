package project

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/shared"
)

type Project struct {
	shared.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Stacks      []string `json:"stacks"`
}
