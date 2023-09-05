package api

import (
	"context"
	"time"

	firestore "cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
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
