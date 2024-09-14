package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of response: %T\n", res)

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Print("Response:\n", string(data))

}
