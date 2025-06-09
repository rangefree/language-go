package main

import (
	"fmt"
	named_import "net/http"
)

func main() {
	fmt.Println("Hello Go standard Library!")
	resp, err := named_import.Get("https://jsonplaceholder.typicode.com/posts/1")
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Erorr: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP responce status:", resp.Status)
}
