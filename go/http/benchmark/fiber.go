package main

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v3"
)

func fiberFramework() framework {
	return framework{
		name: "fiber",
		serve: func(ln net.Listener) func() {
			app := fiber.New()

			// c.SendStatus would put the status text in the body; the other
			// frameworks answer with an empty one.
			app.Get("/ping", func(c fiber.Ctx) error {
				c.Status(fiber.StatusOK)
				return nil
			})

			app.Get("/users/:id", func(c fiber.Ctx) error {
				u := sampleUser
				u.ID = c.Params("id")
				return c.JSON(u)
			})

			app.Post("/users", func(c fiber.Ctx) error {
				var u User
				if err := c.Bind().Body(&u); err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}
				return c.JSON(u)
			})

			go func() {
				cfg := fiber.ListenConfig{DisableStartupMessage: true}
				if err := app.Listener(ln, cfg); err != nil {
					log.Printf("fiber serve: %v", err)
				}
			}()

			return func() {
				if err := app.Shutdown(); err != nil {
					log.Printf("fiber shutdown: %v", err)
				}
			}
		},
	}
}
