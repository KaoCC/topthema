package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Record struct {
	Title string
	Link  string
}

type config struct {
	Token string `json:"OAUTH_TOKEN"`
}

func ReadToken(configFile string) (string, error) {

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("Error when opening config file: ", err)
		return "", err
	}

	var payload config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Println("Error during Unmarshal(): ", err)
		return "", err
	}

	return payload.Token, nil
}

func GetLastTime(filename string) time.Time {

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

func SetLastTime(filename string, last time.Time) {

	err := os.WriteFile(filename, []byte(last.Format(time.RFC3339)), 0644)

	if err != nil {
		log.Println("Fail to write file", err)
	}

}
