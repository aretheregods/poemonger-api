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
	coll := r.DB.Collection(r.Categories)
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
		categories = append(categories, &category)
	}

	return c.Render("Category/list", fiber.Map{
		"Title":      "Poemonger",
		"Subtitle":   "These are categories",
		"Categories": &categories,
	})
}

func (r *APIRoutes) GetCategory(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	poems := []*db.Poem{}
	coll := r.DB.Collection(r.Poetry)
	docs := coll.Where("categories", "array-contains", c.Params("id")).Documents(ctxTimeout)

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return SendBasicError(c, err, fiber.StatusBadGateway)
		}
		poem := db.Poem{}
		err = doc.DataTo(&poem)
		if err != nil {
			return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
		}
		poem.ID = doc.Ref.ID
		poems = append(poems, &poem)
	}

	return c.Render("Category/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a category.",
		"ParentData": fiber.Map{
			"ParentLink":  "/poetry",
			"ParentTitle": "Back to Poetry",
		},
		"Poems": &poems,
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

	category := new(db.Category)
	if err := c.BodyParser(category); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	coll := r.DB.Collection(r.Categories)
	_, err := coll.Doc(category.Name).Set(ctxTimeout, &category)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": fmt.Sprint(category.Name),
		"link": fiber.Map{
			"rel":     "noreferrer",
			"_target": "self",
			"href":    fmt.Sprintf("/%v", r.Categories),
		},
	})
}

func (r *APIRoutes) DeleteCategory(c *fiber.Ctx) error {
	type category struct {
		Name string `json:"name"`
	}
	cat := new(category)
	if err := c.BodyParser(cat); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	coll := r.DB.Collection(r.Categories)
	_, err := coll.Doc(cat.Name).Delete(ctxTimeout)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
