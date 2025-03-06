package urlHandling

import (
	"fmt"
	"net/url"
)

func UrlHandling() {
	fmt.Println("URL Handling")

	myUrl := "https://jsonplaceholder.typicode.com/posts/1?key1=value1&key2=value2"
	fmt.Printf("URL is - %s, \t Type of URL - %T\n", myUrl, myUrl)

	// convert URL into ParseResult
	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error in Parse URL-", err)
		return
	}

	fmt.Printf("Parsed URL is - %s, \t Type of Parsed URL - %T\n", parsedURL, parsedURL)

	// get query parameter in the url
	fmt.Println("Parsed URL Scheme - ", parsedURL.Scheme)
	fmt.Println("Parsed URL Host - ", parsedURL.Host)
	fmt.Println("Parsed URL Path - ", parsedURL.Path)
	fmt.Println("Parsed URL RawQuery - ", parsedURL.RawQuery)

	// Modify Parsed URL Component
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "key2=value3"

	newURL := parsedURL.String()
	fmt.Println("New URL is - ", newURL)
}
