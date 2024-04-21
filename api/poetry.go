package api

import (
	"database/sql"
	"fmt"
	"poemonger/api/db"
	"poemonger/api/with_time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

func (r *APIRoutes) GetAllPoems(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	poems := []*db.Poem{}
	rows, err := r.DB.QueryContext(ctxTimeout, fmt.Sprintf("select * from %s", r.Categories))
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	for rows.Next() {
		if err == iterator.Done {
			break
		}
		if err != nil {
			return SendBasicError(c, err, fiber.StatusBadGateway)
		}
		poem := db.Poem{}
		err = rows.Scan(&poem)
		if err != nil {
			return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
		}

		poems = append(poems, &poem)
	}

	return c.Render("Poem/list", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "These are poems.",
		"Poems":    &poems,
	})
}

func (r *APIRoutes) GetPoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	p := new(db.Poem)
	err := r.DB.QueryRowContext(ctxTimeout, fmt.Sprintf("select * from %v where id =?", r.Poetry), c.Params("id")).Scan(&p)
	if err == sql.ErrNoRows {
		return SendBasicError(c, err, fiber.StatusNotFound)
	}
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.Render("Poem/index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "This is a poem.",
		"PoemID":   c.Params("id"),
		"ParentData": fiber.Map{
			"ParentLink":  "/poetry",
			"ParentTitle": "Back to Poetry",
		},
		"Poem": &p,
	})
}

func (r *APIRoutes) AddPoemForm(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	categories := []*db.Category{}
	rows, err := r.DB.QueryContext(ctxTimeout, fmt.Sprintf("select * from %v", r.Categories))
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadRequest)
	}

	for rows.Next() {
		category := db.Category{}
		err = rows.Scan(&category)
		if err != nil {
			return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
		}
		categories = append(categories, &category)
	}

	return c.Render("Poem/add", fiber.Map{
		"Title":      "Poemonger",
		"Subtitle":   "This is a form to add a poem.",
		"Categories": &categories,
	})
}

func (r *APIRoutes) PostPoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	p := new(db.NewPoem)

	if err := c.BodyParser(&p); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	res, err := r.DB.ExecContext(
		ctxTimeout, 
		fmt.Sprintf("insert or replace into %v (id, title, author, poem, sample_length, categories, release_date) values (NULL, ?, ?, ?, ?, ?, ?)", r.Poetry),
		 p.Title, p.Author, p.Poem, p.SampleLength, p.Categories, p.ReleaseDate)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadRequest)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return SendBasicError(c, err, fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   fmt.Sprint(id),
		"link": fmt.Sprintf("/%v/%v", r.Poetry, id),
	})
}

func (r *APIRoutes) DeletePoem(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	type poem struct {
		ID string `json:"id"`
	}
	p := new(poem)
	if err := c.BodyParser(p); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	_, err := r.DB.ExecContext(ctxTimeout, fmt.Sprintf("delete from %v where id = ?", r.Poetry), p.ID)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
