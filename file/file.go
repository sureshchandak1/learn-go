package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("sample_file.txt") // creating file
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close() // free resourses after completing function

	content := "Hello world by Rakesh"
	_, errors := io.WriteString(file, content+"\n") // writing data in file
	if errors != nil {
		fmt.Println("Error while writing file:", errors)
		return
	}

	fmt.Println("Successfully created file")

	//readFile()

	readFile2() // not use with large file
}

func readFile() {

	file, err := os.Open("sample_file.txt") // open file
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buffer := make([]byte, 1024)

	// Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
	}
}

func readFile2() {

	content, err := os.ReadFile("sample_file.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}

	fmt.Println(string(content))
}
