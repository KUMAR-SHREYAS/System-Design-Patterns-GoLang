package main

import (
	"builder/request"
	"fmt"
)

func main() {
	req1 := request.NewHttpRequestBuilder("https://api.example.com/data").
		Build()

	req2 := request.NewHttpRequestBuilder("https://api.example.com/data").
		Method("POST").
		Body(`{"key":"value"}`).
		Timeout(10000).
		Build()

	req3 := request.NewHttpRequestBuilder("https://api.example.com/data").
		Method("POST").
		AddHeader("Content-Type", "application/json").
		AddQueryParams("version", "1.0").
		Body(`{"key":"value"}`).
		Timeout(10000).
		Build()

	fmt.Println(req1)
	fmt.Println(req2)
	fmt.Println(req3)
}
