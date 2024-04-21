package api

import (
	"fmt"
	"poemonger/api/db"
	withTime "poemonger/api/with_time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

func (r *APIRoutes) GetAllCategories(c *fiber.Ctx) error {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	categories := []*db.Category{}
	rows, err := r.DB.QueryContext(ctxTimeout, fmt.Sprintf("select * from %s;", r.Categories))
	if err != nil {
		return fmt.Errorf("categories err: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err == iterator.Done {
			break
		}
		if err != nil {
			return SendBasicError(c, err, fiber.StatusBadGateway)
		}
		category := db.Category{}
		err = rows.Scan(&category)
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
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	poems := []*db.Poem{}
	rows, err := r.DB.QueryContext(ctxTimeout, fmt.Sprintf("select * from %v, json_each(%v.categories) where json_each.value = ?;", r.Poetry, r.Poetry), c.Params("id"))
	if err != nil {
		return fmt.Errorf("category err: %v", err)
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
	category := new(db.Category)
	if err := c.BodyParser(category); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	res, err := r.DB.ExecContext(ctxTimeout, fmt.Sprintf("insert into %v (id, name, description) values (NULL, ?, ?);", r.Categories), category.Name, category.Description)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return SendBasicError(c, err, fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": fmt.Sprint(id),
		"link": fiber.Map{
			"rel":     "noreferrer",
			"_target": "self",
			"href":    fmt.Sprintf("/%v", r.Categories),
		},
	})
}

func (r *APIRoutes) DeleteCategory(c *fiber.Ctx) error {
	cat := new(db.Category)
	if err := c.BodyParser(cat); err != nil {
		return SendBasicError(c, err, fiber.StatusUnprocessableEntity)
	}

	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	_, err := r.DB.ExecContext(ctxTimeout, fmt.Sprintf("delete from %v where id = ?;", r.Categories), cat.ID)
	if err != nil {
		return SendBasicError(c, err, fiber.StatusBadGateway)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
