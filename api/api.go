package api

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"poemonger/api/db"
)

func InitializeAPI() {
	engine := html.New("pages", ".html")
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
        Views: engine,
    })

	poems := db.InitializeDB("poems")
	api := &APIRoutes{poems}

	app.Get("/poetry/new", api.AddPoemForm)
	app.Get("/poetry/:id", api.GetPoem)
	app.Get("/poetry", api.GetAllPoems)
	app.Post("/poetry", api.PostPoem)
	app.Get("/work/new", api.AddWorkForm)
	app.Get("/work/:id", api.GetWork)
	app.Get("/work", api.GetWorks)
	app.Post("/work", api.PostWork)
	app.Get("/", api.HomePage)

	log.Fatal(app.Listen(":4321"))
}
