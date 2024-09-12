package main

import "fmt"

func main() {

	age := 25
	name := "Rakesh"
	height := 5.86757657

	fmt.Println("Age:", age, "Height:", height, "Name:", name)

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Type of Age: %T\n", age)
	fmt.Printf("Type of Height: %T\n", height)

	fmt.Printf("Name: %s\n", name)

	fmt.Printf("Age: %d, Height: %.2f, Name: %s", age, height, name)

}
