# Go HTTP Client

This repository contains a Go package that provides a flexible and easy-to-use HTTP client for making HTTP requests. The client allows you to set headers, configure request modifiers, and send requests to a given URL.

The idea is to provide you with chainable functions that make it easy to call HTTP requests, especially if you are familiar with the `Fetch API` in the JavaScript world, which I personally love for its simplicity.

## Features

- Set headers: You can easily set custom headers for your HTTP requests.
- Set the base URL: Create a base URL for your client instance so you can shorten the API URL string and reduce the repetition of the URL pattern.
- Request modifiers: You can add request modifiers, such as modifying the request body or performing custom operations on the request before sending it.
- Flexible configuration: The client allows you to configure the base URL and customize the underlying HTTP client.

## Installation

To use this package, you need to have Go installed on your machine. Then, you can install the package using the following command:

```shell
go get github.com/LightningDev/go-req
```

## Usage

Here's an example of how to use the `Client` struct and its methods:

```go
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

```

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://github.com/LightningDev/go-req/blob/main/LICENSE).
