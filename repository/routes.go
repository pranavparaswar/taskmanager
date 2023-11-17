package repository

import "github.com/gofiber/fiber/v2"

func (repo *Repository) SetupRoutes(app *fiber.App) {
	app.Static("/", "./client/public")

	api := app.Group("/api")
	api.Get("/task", repo.GetTask)
	api.Post("/task", repo.CreateTask)

	api.Patch("/task/:id", repo.EditTask)
	api.Delete("/task/:id", repo.DeleteTask)

}
