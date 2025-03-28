package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var podName string

func main() {
	podName = os.Getenv("POD_NAME")
	fmt.Println("Pod Name:", podName)

	router := fiber.New()
	router.Get("/", helloHandler)
	router.Get("/hello/delay", helloDelayHandler)

	err := router.Listen(":5858")
	if err != nil {
		fmt.Printf("listen error: %v", err)
	}
}

func helloHandler(c *fiber.Ctx) error {
	name := c.Query("name", "Guest")
	s := fmt.Sprintf("Welcom %s on the pod %s! (%s)-(%s)", name, podName, c.Context().LocalIP(), c.Context().RemoteIP())
	slog.Info(s)
	return c.Status(http.StatusOK).SendString(s)
}

func helloDelayHandler(c *fiber.Ctx) error {
	delay := rand.Intn(1000)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	name := c.Query("name", "Guest")
	s := fmt.Sprintf("Welcom %s on the pod %s! (%d ms) (%s)-(%s)", name, podName, delay, c.Context().LocalIP(), c.Context().RemoteIP())
	slog.Info(s)
	return c.Status(http.StatusOK).SendString(s)
}
