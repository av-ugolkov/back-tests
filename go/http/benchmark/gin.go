package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginFramework() framework {
	return framework{
		name: "gin",
		serve: func(ln net.Listener) func() {
			gin.SetMode(gin.ReleaseMode)

			// gin.New, not gin.Default: Default installs Logger+Recovery.
			r := gin.New()

			r.GET("/ping", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			r.GET("/users/:id", func(c *gin.Context) {
				u := sampleUser
				u.ID = c.Param("id")
				c.JSON(http.StatusOK, u)
			})

			r.POST("/users", func(c *gin.Context) {
				var u User
				if err := c.ShouldBindJSON(&u); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				c.JSON(http.StatusOK, u)
			})

			return serveHTTP(ln, r)
		},
	}
}
