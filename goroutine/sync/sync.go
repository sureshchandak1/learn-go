package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("workers task complete")
}

func worker(i int, wg *sync.WaitGroup) {

	defer wg.Done() // Signal that this goroutine is done

	fmt.Printf("worker %d started \n", i)
	// some task is happening
	fmt.Printf("worker %d ended \n", i)
}
