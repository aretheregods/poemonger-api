package api

import (
	"fmt"
	"poemonger/api/db"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

func (r *APIRoutes) GetAllCategories(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	categories := []*db.Category{}
	coll := r.DB.Collection("categories")
	docs := coll.Documents(ctxTimeout)

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return SendBasicError(c, err, fiber.StatusBadGateway)
		}
		category := db.Category{}
		err = doc.DataTo(&category)
		if err != nil {
			return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
		}
		category.ID = doc.Ref.ID
		categories = append(categories, &category)
	}

	return c.Render("Category/list", fiber.Map{
		"Title":      "Poemonger",
		"Subtitle":   "These are categories",
		"Categories": &categories,
	})
}

func (r *APIRoutes) GetCategory(c *fiber.Ctx) error {
	return c.Render("Category/index", fiber.Map{
		"Title":      "Poemonger",
		"Subtitle":   "This is a category.",
		"CategoryID": c.Params("id"),
	})
}

func (r *APIRoutes) AddCategoryForm(c *fiber.Ctx) error {
	return c.Render("Category/add", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a form to add a category.",
	})
}

func (r *APIRoutes) PostCategory(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	category := new(db.NewCategory)
	if err := c.BodyParser(category); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	coll := r.DB.Collection("categories")
	res, _, err := coll.Add(ctxTimeout, &category)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": fmt.Sprint(&res.ID),
		"link": fiber.Map{
			"rel":     "noreferrer",
			"_target": "self",
			"href":    "/categories",
		},
	})
}

func (r *APIRoutes) DeleteCategory(c *fiber.Ctx) error {
	type category struct {
		ID string `json:"id"`
	}
	cat := new(category)
	if err := c.BodyParser(cat); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	coll := r.DB.Collection("categories")
	_, err := coll.Doc(cat.ID).Delete(ctxTimeout)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.SendStatus(fiber.StatusNoContent)
}