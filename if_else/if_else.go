package main

import "fmt"

func main() {

	check(10)
	check(2)
	check(5)

	x := 5
	y := 10
	z := 20

	if x > 5 && y > 5 {
		fmt.Println("x:", x, "y:", y)
	} else {
		fmt.Println("either x smaller then 5 nor y")
	}

	if x > 5 && (y > 5 || z < 40) {
		fmt.Println("Hey, How are you?")
	}

}

func check(x int) {

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is smaller than 5")
	}

}
