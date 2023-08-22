package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *mongo.Database
}

func (routes *APIRoutes) HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Poemonger",
	})
}

func (routes *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	return c.Render("Poems/index", fiber.Map{
		"Title": "Poemonger",
		"Subtitle": "These are poems.",
	})
}

func (routes *APIRoutes) GetPoem(c *fiber.Ctx) error {
	return c.Render("Poem/index", fiber.Map{
		"Title": "Poemonger",
		"Subtitle": "This is a poem.",
	})
}

func (routes *APIRoutes) PostPoem(c *fiber.Ctx) error {
	p := new(db.Poem)

	if err := c.BodyParser(p); err != nil {
		return SendBasicError(c, err, 400)
	}

	coll := routes.DB.Collection("poetry")
	res, err := coll.InsertOne(c.Context(), p)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(res.InsertedID),
		"link": fmt.Sprintf("/poetry/%v", res.InsertedID),
	})
}

func (routes *APIRoutes) GetWorks(c *fiber.Ctx) error {
	return c.Render("Works/index", fiber.Map{
		"Title": "Poemonger",
		"Subtitle": "These are collections of poems.",
	})
}

func (routes *APIRoutes) GetWork(c *fiber.Ctx) error {
	return c.Render("Work/index", fiber.Map{
		"Title": "Poemonger",
		"Subtitle": "This is a poetry collection.",
	})
}

func (routes *APIRoutes) PostWork(c *fiber.Ctx) error {
	w := new(db.Work)

	if err := c.BodyParser(w); err != nil {
		return SendBasicError(c, err, 400)
	}

	coll := routes.DB.Collection("works")
	res, err := coll.InsertOne(c.Context(), w)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(res.InsertedID),
		"link": fmt.Sprintf("/work/%v", res.InsertedID),
	})
}
