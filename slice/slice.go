package main

import "fmt"

func main() {

	// Create slice
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("Numbers data type: %T\n", numbers)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Numbers length:", len(numbers))
	fmt.Println("Numbers Capacity:", cap(numbers))

	// Add more data in numbers
	numbers = append(numbers, 6, 7, 8, 9, 10)

	fmt.Println("Numbers:", numbers)
	fmt.Println("Numbers length:", len(numbers))
	fmt.Println("Numbers Capacity:", cap(numbers))

	name := []string{}
	fmt.Println("Name: ", name)
	name = append(name, "Rakesh", "Mohan")
	fmt.Println("Name: ", name)

	// Create alice with custom capacity
	nums := make([]int, 3, 5)

	fmt.Println("Nums: ", nums)
	nums = append(nums, 4)
	fmt.Println("Nums: ", nums)
	nums = append(nums, 98)
	fmt.Println("Nums: ", nums)
	nums = append(nums, 6)
	fmt.Println("Nums: ", nums)

	fmt.Println("Nums length:", len(nums))
	fmt.Println("Nums Capacity:", cap(nums))

	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)
	nums = append(nums, 6)

	fmt.Println("Nums length:", len(nums))
	fmt.Println("Nums Capacity:", cap(nums))

	stock := make([]int, 0)
	fmt.Println("Stock:", stock)

}
