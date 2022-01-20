package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port     string // port without ':'
	MongoURL string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Err with env load:" + err.Error())
		return
	}
	Port = os.Getenv("PORT")
	MongoURL = os.Getenv("MONGO_URL")
}
