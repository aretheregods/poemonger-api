package api

import (
	"log"

	"cloud.google.com/go/firestore"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func InitializeAPI(c *firestore.Client) {
	engine := html.New("pages", ".html")
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Prefork: true,
        Views: engine,
    })

	api := &APIRoutes{c}
	defer closeDBConnection(c)

	app.Get("/categories/new", api.AddCategoryForm)
	app.Get("/categories/:id", api.GetCategory)
	app.Get("/categories", api.GetAllCategories)
	app.Post("/categories", api.PostCategory)
	app.Get("/poetry/new", api.AddPoemForm)
	app.Get("/poetry/:id", api.GetPoem)
	app.Get("/poetry", api.GetAllPoems)
	app.Post("/poetry", api.PostPoem)
	app.Get("/works/new", api.AddWorkForm)
	app.Get("/works/:id", api.GetWork)
	app.Get("/works", api.GetWorks)
	app.Post("/works", api.PostWork)
	app.Get("/", api.HomePage)

	app.Static("/categories*", "./pages/Category/static")
	app.Static("/categories", "./pages/Category/static")
	app.Static("/poetry*", "./pages/Poem/static")
	app.Static("/poetry", "./pages/Poem/static")
	app.Static("/works*", "./pages/Work/static")
	app.Static("/works", "./pages/Work/static")
	app.Static("/", "./pages/static")

	log.Fatal(app.Listen(":4321"))
}

func closeDBConnection(c *firestore.Client) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
