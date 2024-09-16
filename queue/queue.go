package main

import (
	"container/list"
	"fmt"
)

func main() {

	queue := list.New()

	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	for queue.Len() != 0 {

		front := queue.Front()

		fmt.Println(front.Value)

		queue.Remove(front)
	}
}
