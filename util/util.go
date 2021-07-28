package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Record struct {
	Title string
	Link  string
}

func ReadEnv(key string) string {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return os.Getenv(key)
}
