package main

import "fmt"

func main() {

	for i := range 10 {
		fmt.Printf("factorial(%d) = %d\n", i, factorial(i))
	}

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
