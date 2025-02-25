package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type", "Fingerprint"},
	}))
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		// createUserRs := BindFile{
		// 	Name:  "test",
		// 	Email: "test",
		// }

		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		// Save uploaded file
		// file := bindFile.File
		// dst := filepath.Base(file.Filename)
		// if err := c.SaveUploadedFile(file, dst); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("upload file err: %s", err.Error())})
		// 	return
		// }

		c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully."))
	})
	router.Run(":8080")
}
