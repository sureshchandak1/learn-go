package main

import "fmt"

func main() {
	fmt.Println("Learning Array in GoLang")

	// Creating array
	var name [5]string

	name[0] = "Rakesh"
	name[2] = "Mohan"
	fmt.Println(name)

	var numbers = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)

	fmt.Println("Length of Numbers array is:", len(numbers))

	fmt.Println("Value of name at 2 index is:", name[2])

	var price [5]int // store default value 0
	fmt.Println("Price is:", price)

	var product [5]string // store default value ""
	fmt.Println("Product is:", product)
	fmt.Printf("Product is: %q\n", product)

	product[0] = "Milk"
	product[3] = "Book"
	fmt.Printf("Product is: %q\n", product)

}
