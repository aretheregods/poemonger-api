package api

import (
	"fmt"
	"poemonger/api/db"
	"poemonger/api/with_time"

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
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	w := new(db.Work)

	if err := c.BodyParser(w); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	res, err := r.DB.ExecContext(ctxTimeout, fmt.Sprintf("insert or replace into %v ()", r.Works))
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return SendBasicError(c, err, fiber.StatusNotFound)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   fmt.Sprint(id),
		"link": fmt.Sprintf("/%v/%v", r.Works, id),
	})
}