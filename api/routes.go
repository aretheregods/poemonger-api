package api

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *firestore.Client
}

func (routes *APIRoutes) HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "Recent and popular poetry",
	})
}

func (routes *APIRoutes) GetAllCategories(c *fiber.Ctx) error {
	categories := []*db.Category{}
	coll := routes.DB.Collection("categories")
	res := coll.Documents(context.TODO())
	docs, err := res.GetAll()
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	for _, doc := range docs {
		category := db.Category{}
		err := doc.DataTo(&category)
		if err != nil {
			return SendBasicError(c, err, 400)
		}
		category.ID = doc.Ref.ID
		categories = append(categories, &category)
	}

	return c.Render("Category/list", fiber.Map{
		"Title": "Poemonger",
		"Subtitle": "These are categories",
		"Categories": &categories,
	})
}

func (routes *APIRoutes) GetCategory(c *fiber.Ctx) error {
	return c.Render("Category/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a category.",
		"CategoryID": c.Params("id"),
	})
}

func (routes *APIRoutes) AddCategoryForm(c *fiber.Ctx) error {
	return c.Render("Category/add", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a form to add a category.",
	})
}

func (routes *APIRoutes) PostCategory(c *fiber.Ctx) error {
	category := new(db.NewCategory)

	if err := c.BodyParser(category); err != nil {
		return SendBasicError(c, err, 400)
	}

	coll := routes.DB.Collection("categories")
	res, _, err := coll.Add(c.Context(), &category)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(&res.ID),
		"link": fiber.Map{
			"rel": "noreferrer",
			"_target": "self",
			"href": "/categories",
		},
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
	categories := []*db.Category{}
	coll := routes.DB.Collection("categories")
	res := coll.Documents(context.TODO())
	docs, err := res.GetAll()
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	for _, doc := range docs {
		category := &db.Category{}
		err := doc.DataTo(category)
		if err != nil {
			return SendBasicError(c, err, 400)
		}
		categories = append(categories, category)
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
	res, _, err := coll.Add(c.Context(), &p)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(&res.ID),
		"link": fmt.Sprintf("/poetry/%v", &res.ID),
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
	res, _, err := coll.Add(c.Context(), &w)
	if err != nil {
		return SendBasicError(c, err, 400)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(&res.ID),
		"link": fmt.Sprintf("/works/%v", &res.ID),
	})
}
