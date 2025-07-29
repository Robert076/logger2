package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	endpoint := "http://api:8080/"
	body := []byte(`{
		"message": "Hello message from logger service"
	}`)

	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	defer res.Body.Close()

}
