package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var podName string

func main() {
	podName = os.Getenv("POD_NAME")
	fmt.Println("Pod Name:", podName)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.SetTrustedProxies(nil) //show a real address by c.ClientIP()
	router.GET("/", helloHandler)
	router.GET("/hello/delay", helloDelayHandler)

	err := http.ListenAndServe(":5858", router)
	if err != nil {
		fmt.Printf("listen error: %v", err)
	}
}

func helloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	slog.Info(fmt.Sprintf("Welcom %s on the pod %s! (%s)", name, podName, c.ClientIP()))
	c.String(http.StatusOK, "Welcom %s on the pod %s! (%s)", name, podName, c.ClientIP())
}

func helloDelayHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")

	delay := rand.Intn(1000)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	slog.Info(fmt.Sprintf("Welcom %s on the pod %s! (%d ms) (%s)", name, podName, delay, c.ClientIP()))
	c.String(http.StatusOK, "Welcom %s on the pod %s! (%d ms) (%s)", name, podName, delay, c.ClientIP())
}
