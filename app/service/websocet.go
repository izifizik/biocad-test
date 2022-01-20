package service

import (
	"Biocad/app/model"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

var ch = make(chan []byte)

func WriteToWs(tip, fileName string) {
	message := model.Message{
		Type:     tip,
		FileName: fileName,
	}
	data, err := json.Marshal(message)
	if err != nil {
		log.Println("Error with marshal:" + err.Error())
	}
	ch <- data
}

func Write(conn *websocket.Conn) {
	for data := range ch {
		err := conn.WriteMessage(1, data)
		if err != nil {
			log.Println("write message ws: " + err.Error())
			continue
		}
	}
}
