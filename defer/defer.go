package main

import "fmt"

/**
- If any line with defer keyword then this line excute end of funnction
- Multiple defer statements use stack (last-in, first-out [LIFO])
*/
func main() {
	fmt.Println("Start")
	defer fmt.Println(add(10, 10)) // this line excute end of function
	defer fmt.Println("Middle")
	fmt.Println("End")
}

func add(a, b int) int {
	return a + b
}
