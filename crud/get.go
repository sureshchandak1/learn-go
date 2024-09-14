package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CRUD = Create Read Update Delete
func main() {

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
