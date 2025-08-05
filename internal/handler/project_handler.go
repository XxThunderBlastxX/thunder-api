package handler

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/models"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (h *ProjectHandler) CreateProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.Project

		// Parse request body
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		// Create project
		if err := h.service.CreateProject(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

func (h *ProjectHandler) GetProjectByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get project ID from params
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		// Get project from service
		proj, err := h.service.GetProjectByID(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(proj)
	}
}

func (h *ProjectHandler) GetAllProjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get projects from service
		projects, err := h.service.ListProjects()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(projects)
	}
}

func (h *ProjectHandler) UpdateProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get project ID from params
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		var req models.Project

		// Parse request body
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		// Update project
		if err := h.service.UpdateProject(id, &req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

func (h *ProjectHandler) DeleteProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get project ID from params
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		// Delete project
		err := h.service.DeleteProject(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}
}
