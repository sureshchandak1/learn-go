package main

import "fmt"

/**
- If variable name start with small char then this variable use on only this file not allow outside package
ex: var person string = "Rakesh"
- If variable name start with capital char then allow inside and outside both
ex: var Person string = "Mohan"
*/
func main() {

	var name string = "GoLang"
	fmt.Println(name)

	var version = "1.0"
	fmt.Println(version)

	var money int = 67000
	println(money)

	var currency = 3489
	println(currency)

	var dimension float64 = 87.12
	fmt.Println(dimension)

	var decided bool = false
	decided = true
	fmt.Println(decided)

	const pi = 67.12
	// pi = 45 - Not allow
	fmt.Println(pi)

	// Create variable
	person := "Sunil Mohan"
	fmt.Println(person)
}
