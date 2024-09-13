package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")

		if i == 10 {
			fmt.Println()
		}
	}

	counter := 0
	/*for {
		fmt.Print("Infinite Loop")
		counter++
	}*/

	for {
		fmt.Print(counter, " ")
		counter++
		if counter > 5 {
			fmt.Println()
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers {
		fmt.Print(value, " ")

		if index == len(numbers)-1 {
			fmt.Println()
		}
	}

	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Println() // for new line

	data := "abcdefghigklmnopqrstuvwxyz"

	for _, char := range data {
		fmt.Printf("%d-%c, ", char, char)
	}

}
