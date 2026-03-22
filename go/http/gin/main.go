package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.GET("/download", middleware(bindFile))
	router.Run(":8080")
}

func bindFile(c *gin.Context) {
	var bindFile BindFile

	if err := c.ShouldBind(&bindFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File downloaded successfully."))
}

func middleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		next(c)
	}
}
