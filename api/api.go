package api

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *mongo.Database
}

func InitializeAPI()  {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nothing to see here")
	})

	poems := db.InitializeDB("poems")
	api := &APIRoutes{ poems }

	app.Get("/poetry", api.GetAllPoems)
	app.Get("/poetry/:id", api.GetPoem)
	app.Post("/poetry", api.PostPoem)

	app.Listen(":4321")
}

func (routes *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) GetPoem(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) PostPoem(c *fiber.Ctx) error {
	return nil
}