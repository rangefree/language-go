package main

import (
	"fmt"
)

var middleName = "Cane"

func main() {
	// var age int = 0
	// var name string = "John"
	// var name1 = "Doe"

	// no var {!} no type {!} can ONLY be used in functions
	count := 10
	lastName := "Fawn"

	fmt.Println(middleName, lastName, ", count =", count)

	//Default values:
	//	numeric: 0
	//	boolean: false
	//	string: ""
	//	Pointer, slice, map, function, struct: nil

	//Types: PascalCase
	//	structs, interfaces, enums, etc. => UserInfo, CalcAreaa, etc.

	//Variables: snake_case or mixedCase
	//	user_id, first_name, etc.

	//Constants: UPPERCASE

	//Packages: low case and no underscores

	const XEP = "XEP"
	const (
		a = 1
		b = 2
	)

	fmt.Println("Constants:", XEP, ", ", a, ", ", b)

	var f float64 = 22 / 7.0 // one must be float
	fmt.Println(f)

	fmt.Println(22 / 7.0)

	for a := 0; a < 10; a++ {
		if a > 5 {
			fmt.Println("skipping", a)
			break
		}
		//else
		fmt.Println(a)
	}

	//Slice sample:
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		if value%2 == 0 {
			continue
		}
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}

	rows := 8
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 5 {
		i++
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println("in while mode:", i)
		i++
	}

	sum := 0
	for {
		sum += 1
		if sum < 10 {
			continue
		}
		fmt.Println("exiting infinite loop")
		break
	}

}

func printName() {
	firstName := "Michle"
	fmt.Println(firstName)
}
