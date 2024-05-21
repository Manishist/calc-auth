package controllers

import (
	"go-auth/database"
	"go-auth/models"

	"github.com/gofiber/fiber/v2"
)

type ProjectPayload struct {
	UserEmail   string   `json:"user_email"`
	UserName    string   `json:"user_name"`
	ProjectName string   `json:"project_name"`
	History     []string `json:"history"`
}

type UpdateHistoryPayload struct {
	UserEmail   string   `json:"user_email"`
	ProjectName string   `json:"project_name"`
	History     []string `json:"history"`
}

func CreateProject(c *fiber.Ctx) error {
	payload := new(ProjectPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	data := models.Data{
		UserEmail:   payload.UserEmail,
		UserName:    payload.UserName,
		ProjectName: payload.ProjectName,
		History:     payload.History,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create project",
		})
	}

	return c.JSON(data)
}

func GetUserProjects(c *fiber.Ctx) error {
	type UserEmailPayload struct {
		UserEmail string `json:"user_email"`
	}

	payload := new(UserEmailPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var projectNames []string
	if err := database.DB.Model(&models.Data{}).
		Where("user_email = ?", payload.UserEmail).
		Where("deleted_at IS NULL").
		Pluck("project_name", &projectNames).
		Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve projects",
		})
	}

	return c.JSON(fiber.Map{
		"projects": projectNames,
	})
}

func GetProjectHistory(c *fiber.Ctx) error {
	type ProjectHistoryPayload struct {
		UserEmail   string `json:"user_email"`
		ProjectName string `json:"project_name"`
	}

	payload := new(ProjectHistoryPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var project models.Data
	if err := database.DB.Where("user_email = ? AND project_name = ?", payload.UserEmail, payload.ProjectName).First(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve project history",
		})
	}

	return c.JSON(fiber.Map{
		"history": project.History,
	})
}

func UpdateProjectHistory(c *fiber.Ctx) error {
	payload := new(UpdateHistoryPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var project models.Data
	if err := database.DB.Where("user_email = ? AND project_name = ?", payload.UserEmail, payload.ProjectName).First(&project).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	historySet := make(map[string]bool)
	for _, entry := range project.History {
		historySet[entry] = true
	}

	for _, newEntry := range payload.History {
		if !historySet[newEntry] {
			project.History = append(project.History, newEntry)
			historySet[newEntry] = true
		}
	}

	if err := database.DB.Save(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update project history",
		})
	}

	return c.JSON(project)
}
