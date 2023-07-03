package main

import (
	"log"
	"net/http"

	client "github.com/LightningDev/go-req"
)

func main() {
	get()
	post()
}

// GET request
func get() {
	client := client.New("https://animechan.xyz")
	resp, err := client.Fetch("GET", "/api/random")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response: " + string(resp) + "\n")
}

// POST request
func post() {
	client := client.New("https://api-ssl.bitly.com")
	data := `{ "long_url": "https://github.com/LightningDev/go-req", "domain": "bit.ly" }`

	header := &http.Header{}
	header.Set("Authorization", "Bearer {YOUR_BITLY_TOKEN}")
	header.Set("Content-Type", "application/json")

	resp, err := client.
		SetHeadersFromObject(header).
		SetBody(data).
		Fetch("POST", "/v4/shorten")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response: " + string(resp) + "\n")
}
