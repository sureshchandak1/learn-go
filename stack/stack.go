package main

import (
	"container/list"
	"fmt"
)

func main() {

	stack := list.New()
	fmt.Printf("Stack type: %T\n", stack)

	for i := 1; i <= 10; i++ {
		stack.PushBack(i)
	}

	for stack.Back() != nil {

		back := stack.Back()

		fmt.Println(back.Value)

		stack.Remove(back)
	}

	stack = list.New()

	stack.PushBack("a")
	stack.PushBack("b")
	stack.PushBack("c")
	stack.PushBack("d")

	for stack.Len() != 0 {

		back := stack.Back()

		fmt.Println(back.Value)

		stack.Remove(back)
	}

}
