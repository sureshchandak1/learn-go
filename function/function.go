package main

import "fmt"

func main() {
	fmt.Println("We are learning function in GoLang")

	simpleFunction()

	ans := add(3, 4)
	fmt.Println(ans)
	fmt.Println(addValue(5, 5))
}

func simpleFunction() {
	fmt.Println("Simple function")
}

func sum(a int, b int) int {
	return a + b
}

func add(a, b int) int {
	return a + b
}

func addValue(a, b int) (result int) {
	result = a + b
	return
}
