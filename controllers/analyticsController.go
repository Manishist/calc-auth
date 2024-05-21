package controllers

import (
	"go-auth/database"
	"go-auth/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	// Retrieve all users from the database
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	// Prepare response data for each user
	var responseData []fiber.Map
	for _, user := range users {
		// Calculate total time consumption
		totalTimeConsumed := user.TotalTimeConsumed

		// Calculate total time today
		totalTimeToday := user.TotalTimeToday

		// Calculate logged in days in the last 7 days
		loggedInDaysLast7Days := user.LoggedInDaysLast7Days

		// Prepare user data for response
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
