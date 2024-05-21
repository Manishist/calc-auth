package routes

import (
	"go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Post("/api/project", controllers.CreateProject)
	app.Post("/api/user-projects", controllers.GetUserProjects)
	app.Post("/api/project-history", controllers.GetProjectHistory)
	app.Put("/api/update-project-history", controllers.UpdateProjectHistory)

	app.Post("/api/user-info", controllers.GetUserInfo)
}
