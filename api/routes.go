package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *mongo.Database
}

func (routes *APIRoutes) HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "Recent and popular poetry",
	})
}

func (routes *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	return c.Render("Poem/list", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "These are poems.",
	})
}

func (routes *APIRoutes) GetPoem(c *fiber.Ctx) error {
	return c.Render("Poem/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a poem.",
		"PoemID": c.Params("id"),
	})
}

func (routes *APIRoutes) AddPoemForm(c *fiber.Ctx) error {
	categories := []db.Category{}
	coll := routes.DB.Collection("categories")
	res, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return SendBasicError(c, err, 400)
	}
	err = res.All(context.TODO(), &categories)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Render("Poem/add", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a form to add a poem.",
		"Categories": &categories,
	})
}

func (routes *APIRoutes) PostPoem(c *fiber.Ctx) error {
	p := new(db.NewPoem)

	if err := c.BodyParser(p); err != nil {
		return SendBasicError(c, err, 400)
	}

	coll := routes.DB.Collection("poetry")
	res, err := coll.InsertOne(c.Context(), &p)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(res.InsertedID),
		"link": fmt.Sprintf("/poetry/%v", &res.InsertedID),
	})
}

func (routes *APIRoutes) GetWorks(c *fiber.Ctx) error {
	return c.Render("Work/list", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "These are collections of poems.",
	})
}

func (routes *APIRoutes) GetWork(c *fiber.Ctx) error {
	return c.Render("Work/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a poetry collection.",
		"WorkID": c.Params("id"),
	})
}

func (routes *APIRoutes) AddWorkForm(c *fiber.Ctx) error {
	return c.Render("Work/add", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a form to add a work.",
	})
}

func (routes *APIRoutes) PostWork(c *fiber.Ctx) error {
	w := new(db.Work)

	if err := c.BodyParser(w); err != nil {
		return SendBasicError(c, err, 400)
	}

	coll := routes.DB.Collection("works")
	res, err := coll.InsertOne(c.Context(), &w)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(res.InsertedID),
		"link": fmt.Sprintf("/work/%v", &res.InsertedID),
	})
}
