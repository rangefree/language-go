package main

import "fmt"

func main() {
	message := "Hello World!"
	fmt.Println("message:", message)

	//NOTE: range operate on the COPY!!!!
	for i, v := range message {
		fmt.Printf("message[%d] = '%c' (%d)\n", i, v, v)
	}
}
