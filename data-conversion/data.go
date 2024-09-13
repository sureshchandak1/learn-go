package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 42
	fmt.Println("Number is:", num)
	fmt.Printf("Type of num is %T\n", num)

	//num = num + 1.23

	// Type conversion
	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is:", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	// Convert int to string
	str := strconv.Itoa(num)
	fmt.Println("str is:", str)
	fmt.Printf("Type of str is %T\n", str)

	number := "1234"
	fmt.Printf("Type of number is %T\n", number)
	// Convert string to int
	num, _ = strconv.Atoi(number)
	fmt.Println("Number is:", number)
	fmt.Println("num is:", num)
	fmt.Printf("Type of num is %T\n", num)

	numString := "3.14"
	// Convert string to float
	numFloat, _ := strconv.ParseFloat(numString, 64)
	fmt.Println("numFloat is:", numFloat)
	fmt.Printf("Type of numFloat is %T\n", numFloat)

}
