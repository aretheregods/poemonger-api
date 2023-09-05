package api

import (
	"fmt"
	"poemonger/api/db"

	"github.com/gofiber/fiber/v2"
)

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