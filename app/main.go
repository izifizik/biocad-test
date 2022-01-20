package main

import (
	"Biocad/app/config"
	"Biocad/app/database"
	"Biocad/app/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	config.Load()
	database.Connect()

	app := gin.Default()
	gin.SetMode(gin.DebugMode)
	app.LoadHTMLGlob("app/template/*")

	app.GET("/ws", handlers.GetPage)
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})

	app.POST("/upload", handlers.FileUpload)
	app.POST("/delete", handlers.FileDelete)

	app.StaticFS("/file", http.Dir("app/public"))

	err := app.Run("0.0.0.0:8080")
	if err != nil {
		log.Println("Backend do not start with error: " + err.Error())
	}
}
