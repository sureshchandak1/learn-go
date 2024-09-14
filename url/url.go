package main

import (
	"fmt"
	"net/url"
)

/*
			APIs
	  URL - Uniform Resource Locators
*/
func main() {
	fmt.Println("Learning URl")

	getURL := "https://example.com:8080/path/to/resourse?key1=value1&key2=value2"

	parsedURL, err := url.Parse(getURL)
	if err != nil {
		fmt.Println("Can't parse URL", err)
	}

	fmt.Printf("Type of URL: %T\n", parsedURL)

	fmt.Println("Scheme", parsedURL.Scheme)
	fmt.Println("Host", parsedURL.Host)
	fmt.Println("Path", parsedURL.Path)
	fmt.Println("RawQuery", parsedURL.RawQuery)

	// Modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=rakesh"

	// Constructing a URL string grom a URL object
	newUrl := parsedURL.String()

	fmt.Println("new URL:", newUrl)

}
