package main

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func main() {
	file, err := os.CreateTemp("", "testTempFile-")
	Check(err)

	fmt.Printf("Temp file %s was created.\n", file.Name())

	defer func() {
		os.Remove(file.Name())
		fmt.Printf("Temp file %s was removed.\n", file.Name())
	}()

	defer func() {
		file.Close()
		fmt.Printf("Temp file %s was closed.\n", file.Name())
	}()
	// you can do the same with tempDirectories: os.MkdirTemp("aaa", "ttt"), os.RemoveAll("aaa")
}
