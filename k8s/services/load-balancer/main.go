package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", helloHandler)
	router.GET("/hello", helloDelayHandler)
	router.GET("/delay", delayHandler)

	err := http.ListenAndServe(":5858", router)
	if err != nil {
		fmt.Printf("listen error: %v", err)
	}
}

func helloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	slog.Info(fmt.Sprintf("helloHandler: %s", name))
	c.String(http.StatusOK, "Hello %s!", name)
}

func helloDelayHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")

	delay := rand.Intn(1000)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	slog.Info(fmt.Sprintf("helloDelayHandler: %s (%d ms)", name, delay))

	c.String(http.StatusOK, "Hello %s! (%d ms)", name, delay)
}

func delayHandler(c *gin.Context) {
	d := c.DefaultQuery("d", "0")

	delay, _ := strconv.Atoi(d)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	slog.Info(fmt.Sprintf("delayHandler: %d ms", delay))

	c.String(http.StatusOK, "%d ms", delay)
}
