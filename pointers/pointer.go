package main

import "fmt"

func main() {

	num := 2

	// Create pointer to point memory address of num variable
	var ptr *int = &num

	fmt.Println("Memory Address of num:", ptr)
	fmt.Println("num contain value:", num)

	fmt.Println("Data contains through Pointer:", *ptr)

	var pointer *int // default value nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	fmt.Println("Value contains:", value)
	modifyValueByReference(&value)
	fmt.Println("Value contains:", value)
}

func modifyValueByReference(num *int) {
	*num = *num + 20
}
