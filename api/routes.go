package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type APIRoutes struct {
	DB         *sql.DB
	Categories string
	Poetry     string
	Works      string
	SignUp     string
	Login      string
	Logout     string
	Reset      string
	Delete     string
}

func (r *APIRoutes) HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Poemonger",
		"Subtitle": "Recent and popular poetry",
	})
}
