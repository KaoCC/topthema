package util

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Record struct {
	Title string
	Link  string
}

const envFile string = "app.env"

func ReadEnv(key string) string {

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return os.Getenv(key)
}

const filename string = "last.txt"

func GetLastTime() time.Time {

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Fail to read file", err)
		return time.Time{}
	}

	last, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		log.Println("Fail to parse tile data", err)
		return time.Time{}
	}

	return last
}

func SetLastTime(last time.Time) {

	err := os.WriteFile(filename, []byte(last.Format(time.RFC3339)), 0644)

	if err != nil {
		log.Println("Fail to write file", err)
	}

}
