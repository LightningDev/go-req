package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	BaseURL      *url.URL
	Client       *http.Client
	ReqModifiers []func(*http.Request)
}

func New(baseURL string) *Client {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println(fmt.Sprintf("Error parsing URL: %s", baseURL), err)
		return nil
	}

	return &Client{
		BaseURL:      parsedURL,
		Client:       &http.Client{},
		ReqModifiers: make([]func(*http.Request), 0),
	}
}

func (c *Client) SetHeader(key, value string) *Client {
	headers := func(req *http.Request) {
		req.Header.Add(key, value)
	}

	c.ReqModifiers = append(c.ReqModifiers, headers)
	return c
}

func (c *Client) SetHeadersFromObject(headers *http.Header) *Client {
	for k, v := range *headers {
		c.SetHeader(k, v[0])
	}
	return c
}

func (c *Client) SetBody(body string) *Client {
	contentType := http.DetectContentType([]byte(body))

	modifier := func(req *http.Request) {
		req.Body = io.NopCloser(strings.NewReader(body))
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", contentType)
		}
	}

	c.ReqModifiers = append(c.ReqModifiers, modifier)
	return c
}

func (c *Client) sendRequest(method, endpoint string) (*http.Response, error) {
	url, err := c.BaseURL.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		log.Println("Request Error:", err)
		return nil, err
	}

	for _, reqModifier := range c.ReqModifiers {
		reqModifier(req)
	}
	c.ReqModifiers = make([]func(*http.Request), 0)

	return c.Client.Do(req)
}

func (c *Client) Fetch(method string, path string) ([]byte, error) {
	log.Println("Sending request to " + path)

	var response *http.Response = nil
	var responseError error

	response, responseError = c.sendRequest(method, path)

	if responseError != nil {
		log.Println("Response Error:", responseError)
		return nil, responseError
	}
	defer response.Body.Close()

	log.Println("Response Status:", response.Status)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Response Body Error:", responseError)
		return nil, err
	}

	return responseBody, nil
}
