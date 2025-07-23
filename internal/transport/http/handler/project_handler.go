package handler

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/domain/project"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	service project.Service
}

func NewProjectHandler(service project.Service) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (h *ProjectHandler) CreateProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation for creating a project
		return nil
	}
}

func (h *ProjectHandler) GetProjectByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation for getting a project by ID
		return nil
	}
}

func (h *ProjectHandler) GetAllProjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation for getting all projects
		return nil
	}
}

func (h *ProjectHandler) UpdateProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation for updating a project
		return nil
	}
}

func (h *ProjectHandler) DeleteProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation for deleting a project
		return nil
	}
}
