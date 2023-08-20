package api

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *mongo.Database
}

func InitializeAPI() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Nothing to see here!")
	})

	poems := db.InitializeDB("poems")
	api := &APIRoutes{poems}

	app.Get("/poetry/:id", api.GetPoem)
	app.Get("/poetry", api.GetAllPoems)
	app.Post("/poetry", api.PostPoem)
	app.Get("/work/:id", api.GetWork)
	app.Get("/work", api.GetWorks)
	app.Post("/work", api.PostWork)

	app.Listen(":4321")
}

func (routes *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) GetPoem(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) PostPoem(c *fiber.Ctx) error {
	p := new(db.Poem)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	coll := routes.DB.Collection("poetry")
	_, err := coll.InsertOne(c.Context(), p)
	if err != nil {
		return err
	}

	return c.SendStatus(200)
}

func (routes *APIRoutes) GetWorks(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) GetWork(c *fiber.Ctx) error {
	return nil
}

func (routes *APIRoutes) PostWork(c *fiber.Ctx) error {
	return nil
}
