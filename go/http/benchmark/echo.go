package main

import (
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

func echoFramework() framework {
	return framework{
		name: "echo",
		serve: func(ln net.Listener) func() {
			e := echo.New()
			e.HideBanner = true
			e.HidePort = true

			e.GET("/ping", func(c echo.Context) error {
				return c.NoContent(http.StatusOK)
			})

			e.GET("/users/:id", func(c echo.Context) error {
				u := sampleUser
				u.ID = c.Param("id")
				return c.JSON(http.StatusOK, u)
			})

			e.POST("/users", func(c echo.Context) error {
				var u User
				if err := c.Bind(&u); err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}
				return c.JSON(http.StatusOK, u)
			})

			return serveHTTP(ln, e)
		},
	}
}
