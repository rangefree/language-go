package main

import (
	"fmt"
)

func main() {
	sequence := adder() //adder() called once! anf sequence holds returned function (!)
	// in such case, i captured by closure and initialized once!!!
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder() //adder() called once! anf sequence holds returned function (!)
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	subtractor := func() func(int) int {
		countdown := 99 //part of the closure! called once
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	for i := range 100 {
		res := subtractor(i)
		fmt.Printf("(%d, %d) ", i, res)
		if res < 0 {
			fmt.Println("Subtractor reached negative value ", res, ") at i =", i)
			break
		}
	}
}

func adder() func() int {
	i := 0
	fmt.Println("Current value of i =", i)
	return func() int {
		i++
		fmt.Println("New value of i =", i)
		return i
	}
}
