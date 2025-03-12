package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BindFile struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func main() {
	router := echo.New()
	router.Static("/", "./public")
	router.GET("/download", echo.HandlerFunc(bindFile), middleware)
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func bindFile(c echo.Context) error {
	var bindFile BindFile

	if err := c.Bind(&bindFile); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully."))
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("some code")
		return next(c)
	}
}
