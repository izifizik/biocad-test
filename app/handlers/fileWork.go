package handlers

import (
	"Biocad/app/database"
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

	correctFileName, err := service.UploadFile(filename, file)
	if err != nil {
		log.Println("Upload file: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect file",
		})
		return
	}

	err = service.InsertByFileName(correctFileName)
	if err != nil {
		log.Println("Mongo insert error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with db",
		})
		return
	}
	service.WriteToWs("ADD", correctFileName)

	filepath := "http://localhost:8080/file/" + correctFileName
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

	err = database.DeleteFile(jsonInput.FileName)
	if err != nil {
		log.Println("Mongo delete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with db",
		})
		return
	}
	service.WriteToWs("DELETE", jsonInput.FileName)
}
