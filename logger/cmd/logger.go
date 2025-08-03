package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	endpoint := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	postUrl := fmt.Sprintf("%s:%s", endpoint, port)

	body := []byte(`{
		"message": "Hello message from logger service"
	}`)

	r, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	r.Header.Add("Content-Type", "application/json")

	for i := 0; i < 10; i++ {
		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		time.Sleep(2 * time.Second)
		res.Body.Close()
	}
}
