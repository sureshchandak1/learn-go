package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello()
	go sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(1000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello, world!")
	time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Rakesh :)")
	time.Sleep(1000 * time.Millisecond) // Simulating some work
}
