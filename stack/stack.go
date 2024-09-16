package main

import (
	"container/list"
	"fmt"
)

func main() {

	stack := list.New()

	for i := 1; i <= 10; i++ {
		stack.PushBack(i)
	}

	for stack.Back() != nil {

		back := stack.Back()

		fmt.Println(back.Value)

		stack.Remove(back)
	}

}
