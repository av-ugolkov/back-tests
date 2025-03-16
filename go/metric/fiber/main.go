package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	router := fiber.New()
	router.Get("/hello", handlerHello)
	router.Get("/hello/:name", handlerHello)
	router.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	router.Listen(":3000")
}

func handlerHello(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		name = "Unknown"
	}
	msg := fmt.Sprintf("Hello, %s!", name)
	return c.SendString(msg)
}
