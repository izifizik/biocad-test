package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func FileUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("c.Req.FromFile" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Smth wrong in get file on server: " + err.Error(),
		})
		return
	}

	filename := header.Filename
	out, err := os.Create("app/public/" + filename)
	if err != nil {
		log.Println("os.Create" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Can't create file : " + err.Error(),
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("io.Copy" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "internal server error: " + err.Error(),
		})
		return
	}

	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}

func FileDelete(c *gin.Context) {
	jsonInput := struct {
		FileName string `json:"fileName"`
	}{}

	err := c.ShouldBindJSON(&jsonInput)
	if err != nil {
		log.Println("Should bind json " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = os.Remove("app/public/" + jsonInput.FileName)
	if err != nil {
		log.Println("os.Remove " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file",
		})
		return
	}

	filepath := "http://localhost:8080/file/" + jsonInput.FileName
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}
