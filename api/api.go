package api

import (
	"context"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeAPI(db *mongo.Database, c *mongo.Client) {
	engine := html.New("pages", ".html")
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Prefork: true,
        Views: engine,
    })

	api := &APIRoutes{db}
	defer closeDBConnection(c)

	app.Get("/poetry/new", api.AddPoemForm)
	app.Get("/poetry/:id", api.GetPoem)
	app.Get("/poetry", api.GetAllPoems)
	app.Post("/poetry", api.PostPoem)
	app.Get("/work/new", api.AddWorkForm)
	app.Get("/work/:id", api.GetWork)
	app.Get("/work", api.GetWorks)
	app.Post("/work", api.PostWork)
	app.Get("/", api.HomePage)

	app.Static("/poetry", "./pages/Poem/static")
	app.Static("/work", "./pages/Work/static")
	app.Static("/", "./pages/static")

	log.Fatal(app.Listen(":4321"))
}

func closeDBConnection(c *mongo.Client) {
	if err := c.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
