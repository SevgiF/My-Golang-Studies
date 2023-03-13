package main

import (
	"log"

	c "github.com/sfidann/urun_yonetim_elastic/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	api.Post("/product/post", c.PostProduct)
	api.Get("/product/search", c.Search)
	api.Get("/product/filter", c.Filter)

	err := app.Listen(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
