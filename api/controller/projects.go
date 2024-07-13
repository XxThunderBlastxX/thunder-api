package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/internal/db/gen/projectDb"
	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/model"
)

type ProjectsController struct {
	ProjectsService domain.ProjectsService
}

func (p *ProjectsController) AddProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var proj projectDb.Project

		err := c.BodyParser(&proj)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   err.Error(),
			})
		}

		err = p.ProjectsService.CreateProject(&proj)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Success: true,
			Data:    &model.SuccessResponse{Message: "Project added successfully"},
		})
	}
}

func (p *ProjectsController) ListProjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		projects, err := p.ProjectsService.ListProjects()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.ProjectsResponse]{
			Success: true,
			Data:    &model.ProjectsResponse{Projects: projects},
		})
	}
}

func (p *ProjectsController) RemoveProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Query("id", "0"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   "Invalid ID format",
			})
		}

		name := c.Query("name", "")

		if name != "" {
			err := p.ProjectsService.RemoveProjectByName(name)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
					Success: false,
					Error:   err.Error(),
				})
			}
		} else {
			err := p.ProjectsService.RemoveProjectById(int32(id))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
					Success: false,
					Error:   err.Error(),
				})
			}
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Success: true,
			Data:    &model.SuccessResponse{Message: "Project removed successfully"},
		})
	}
}
