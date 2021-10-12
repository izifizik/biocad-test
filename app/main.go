package main

import (
	"Biocad/app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.GetPage)
	//http.HandleFunc("/add", handlers.AddFile)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
