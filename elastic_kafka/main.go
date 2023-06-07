package main

import (
	"log"

	c "github.com/SevgiF/elastic_kafka/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	api.Post("/post", c.Post)

	err := app.Listen(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
