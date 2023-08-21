package api

import (
	"github.com/gofiber/fiber/v2"

	"poemonger/api/db"
)

func InitializeAPI() {
	app := fiber.New()

	poems := db.InitializeDB("poems")
	api := &APIRoutes{poems}

	app.Get("/poetry/:id", api.GetPoem)
	app.Get("/poetry", api.GetAllPoems)
	app.Post("/poetry", api.PostPoem)
	app.Get("/work/:id", api.GetWork)
	app.Get("/work", api.GetWorks)
	app.Post("/work", api.PostWork)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nothing to see here!")
	})

	app.Listen(":4321")
}
