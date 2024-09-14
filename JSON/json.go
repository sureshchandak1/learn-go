package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	person := Person{Name: "John", Age: 40, IsAdult: true}
	fmt.Println(person)

	// Convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling:", err)
		return
	}
	fmt.Println(string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling:", err)
		return
	}
	fmt.Println(personData)

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var todo Todo
	json.Unmarshal(data, &todo)

	fmt.Print(todo)
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
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
