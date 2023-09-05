package api

import (
	"fmt"
	"poemonger/api/db"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

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