package webrequest

import (
	"fmt"
	"io"
	"net/http"
)

func WebRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Getting Error in Web Request-", err)
		return
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Read Response-", err)
		return
	}

	originalData := string(data)

	fmt.Println("URL Response is -", originalData)

}
