package main

import (
	"bytes"
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

	s := "12121211122"
	removeLast := s[0 : len(s)-1]
	first3 := s[0:3]
	last3 := s[len(s)-3:]
	last1 := s[len(s)-1:]

	fmt.Println(s)
	fmt.Println("Remove Last:", removeLast)
	fmt.Println("First 3:", first3)
	fmt.Println("Last 3:", last3)
	fmt.Println("Last 1:", last1)

	fmt.Printf("Last 1 type: %T\n", last1)

	var buffer bytes.Buffer
	buffer.WriteString("a")
	buffer.WriteByte(3)
	fmt.Println("Buffer string:", buffer.String())

}
