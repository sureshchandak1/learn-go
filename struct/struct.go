package main

import "fmt"

func main() {

	var person1 Person

	fmt.Println("Person1:", person1) // print default value

	person1.FirstName = "Shree"
	person1.LastName = "Ram"
	person1.Age = 1000000000

	fmt.Println("Person1:", person1)

	person2 := Person{
		FirstName: "Aakash",
		LastName:  "Singh",
		Age:       26,
	}
	fmt.Println("Person2:", person2)

	// new keyword
	var person3 = new(Person)
	person3.FirstName = "Simran"
	person3.LastName = "Agarwal"
	person3.Age = 30
	fmt.Println("Person3:", person3)

	var emp1 Employee
	emp1.Person = Person{
		FirstName: "Aakash",
		LastName:  "Singh",
		Age:       26,
	}
	emp1.Contact.Email = "contact@gmail.com"
	emp1.Contact.Phone = "8654872689"
	emp1.Address = Address{
		House: 765,
		Area:  "Ranchi",
		State: "Jharkhand",
	}

	fmt.Println("Employee1:", emp1)

}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Employee struct {
	Person  Person
	Contact Contact
	Address Address
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}
