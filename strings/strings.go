package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("two count:", count)

	str = "            Hello,           Go!            "
	trimed := strings.TrimSpace(str)
	fmt.Println("Trimed value:", trimed)

	str1 := "Jay Shree"
	str2 := "Ram"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Join strings:", result)

}
