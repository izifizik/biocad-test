package handlers

import (
	"Biocad/app/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func FileUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("c.Req.FromFile" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Smth wrong in get file on server",
		})
		return
	}

	filename := header.Filename

	err = service.UploadFile(filename, file)
	if err != nil {
		log.Println("Upload file: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Smth wrong in get file on server",
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
