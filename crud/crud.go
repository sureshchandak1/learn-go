package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CRUD = Create Read Update Delete
func main() {

	performGetRequest()

	performPostRequest()

	performUpdateRequest()

}

func performGetRequest() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response:", res.Status)
		return
	}

	// Read the response body
	/*
		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Print("Response:\n", string(data))
	*/

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Todo:", todo)

}

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Rakesh Kumar TODO",
		Completed: true,
	}

	// Convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	jsonString := string(jsonData) // Convert json data to string

	jsonReader := strings.NewReader(jsonString) // Convert string data to reader

	// Send POST request
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	data, _ := io.ReadAll(res.Body)
	fmt.Print("Response:\n", string(data))
	fmt.Println()
}

func performUpdateRequest() {
	todo := Todo{
		UserId:    235345,
		Title:     "Rakesh Kumar & Mohanlal",
		Completed: false,
	}

	// Convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	jsonString := string(jsonData) // Convert json data to string

	jsonReader := strings.NewReader(jsonString) // Convert string data to reader

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	defer req.Body.Close()

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	data, _ := io.ReadAll(res.Body)
	fmt.Print("Response:\n", string(data))
	fmt.Println()

}

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*
{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
}
*/
