package main

import "fmt"

func main() {
	fmt.Println("Started Error Handling in the functions")

	ans := divide(10, 4)
	fmt.Println(ans)

	ans = divide(10, 0)
	fmt.Println(ans)

	result, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// underscore ( _ ) if you want to ignore
	ans, _ = div(10, 0)
	fmt.Println(ans)
}

func divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func divValue(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Denominator must not be zero"
	}
	return a / b, ""
}
