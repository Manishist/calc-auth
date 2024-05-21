package controllers

import (
	"go-auth/database"
	"go-auth/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}
	var responseData []fiber.Map
	for _, user := range users {
		totalTimeConsumed := user.TotalTimeConsumed

		totalTimeToday := user.TotalTimeToday

		loggedInDaysLast7Days := user.LoggedInDaysLast7Days

		userData := fiber.Map{
			"name":                   user.Name,
			"total_time_consumption": totalTimeConsumed,
			"total_time_today":       totalTimeToday,
			"logged_in_last_7_days":  loggedInDaysLast7Days,
		}
		responseData = append(responseData, userData)
	}

	return c.JSON(responseData)
}
