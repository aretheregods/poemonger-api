package api

import (
	"context"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"

	"poemonger/api/db"
)

type APIRoutes struct {
	DB *firestore.Client
}

func WithTimeout(timeoutLength int) (context.Context, context.CancelFunc) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeoutLength))

	return ctxTimeout, cancel
}

func (r *APIRoutes) HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "Recent and popular poetry",
	})
}

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

func (r *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	poems := []*db.Poem{}
	coll := r.DB.Collection("poetry")
	docs := coll.Documents(ctxTimeout)

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

	return c.Render("Poem/list", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "These are poems.",
		"Poems":    poems,
	})
}

func (r *APIRoutes) GetPoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	p := new(db.Poem)
	coll := r.DB.Collection("poetry")
	doc, err := coll.Doc(c.Params("id")).Get(ctxTimeout)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	err = doc.DataTo(&p)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}
	p.ID = doc.Ref.ID

	return c.Render("Poem/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a poem.",
		"PoemID":   c.Params("id"),
		"Poem":     p,
	})
}

func (r *APIRoutes) AddPoemForm(c *fiber.Ctx) error {
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

	return c.Render("Poem/add", fiber.Map{
		"Title":      "Poemonger",
		"Subtitle":   "This is a form to add a poem.",
		"Categories": &categories,
	})
}

func (r *APIRoutes) PostPoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	p := new(db.NewPoem)

	if err := c.BodyParser(&p); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	coll := r.DB.Collection("poetry")
	res, _, err := coll.Add(ctxTimeout, &p)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   fmt.Sprint(&res.ID),
		"link": fmt.Sprintf("/poetry/%v", res.ID),
	})
}

func (r *APIRoutes) DeletePoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	type poem struct {
		ID string `json:"id"`
	}
	p := new(poem)
	if err := c.BodyParser(p); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	coll := r.DB.Collection("poetry")
	_, err := coll.Doc(p.ID).Delete(ctxTimeout)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (r *APIRoutes) GetWorks(c *fiber.Ctx) error {
	return c.Render("Work/list", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "These are collections of poems.",
	})
}

func (r *APIRoutes) GetWork(c *fiber.Ctx) error {
	return c.Render("Work/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a poetry collection.",
		"WorkID":   c.Params("id"),
	})
}

func (r *APIRoutes) AddWorkForm(c *fiber.Ctx) error {
	return c.Render("Work/add", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a form to add a work.",
	})
}

func (r *APIRoutes) PostWork(c *fiber.Ctx) error {
	ctxTimeout, cancel := WithTimeout(10)
	defer cancel()

	w := new(db.Work)

	if err := c.BodyParser(w); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	coll := r.DB.Collection("works")
	res, _, err := coll.Add(ctxTimeout, &w)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(&res.ID),
		"link": fmt.Sprintf("/works/%v", &res.ID),
	})
}
