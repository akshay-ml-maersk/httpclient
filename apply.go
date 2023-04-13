package main

import (
	"encoding/json"
	"fmt"

	pensieve "golan/library"
)

func main() {
	client := pensieve.NewPensieveClient("https://httpbin.org")

	requestBody := map[string]string{"hello": "world"}
	responseBody, err := client.SendRequest("POST", "/post", requestBody)
	if err != nil {
		fmt.Printf("error retrieving data: %s\n", err)
		return
	}

	jsonData, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Printf("error marshaling response body: %s\n", err)
		return
	}

	fmt.Printf("data: %s\n", string(jsonData))
}
