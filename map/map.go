package main

import "fmt"

func main() {

	// Create map: key string type and value int type
	studentGrades := make(map[string]int)

	studentGrades["Mohan"] = 34
	studentGrades["Rakesh"] = 45
	studentGrades["Shiv"] = 100
	studentGrades["Hanuman"] = 100

	fmt.Println(studentGrades)
	fmt.Println(studentGrades["Mohan"])

	studentGrades["Mohan"] = 50 // Update value

	fmt.Println(studentGrades["Mohan"])

	delete(studentGrades, "Mohan")
	fmt.Println(studentGrades)

	// key exists or not
	value, exists := studentGrades["Mohan"]
	fmt.Println("Grades of Mohan:", value)
	fmt.Println("Mohan exists:", exists)

	for key, value := range studentGrades {
		fmt.Println("Key:", key, "Value:", value)
	}

}
