package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// init function(s) of the package will be called by GO before main() !
// init function(s) will be called in order they are specified in the program
func init() {
	fmt.Println("1. init() entered")
}

func init() {
	fmt.Println("2. init() entered")
}

func init() {
	fmt.Println("3. init() entered")
}

func main() {
	fmt.Println("function main() entered")
	//func NAME(parameters list) return type(s) {
	// code block
	// return value
	//}

	fmt.Println(add(1, 2))
	greet := func(name string) {
		fmt.Println("Hello", name, "from anonymous function!")
	}
	greet("Dennis")

	operation := add

	fmt.Println("call add() as operation(1,2):", operation(1, 2))

	fmt.Println("applyOperation(3,5,add) =", applyOperation(3, 5, add))

	multBy2 := createMultiplier(2)
	fmt.Println("Multiply 3 by 2 =", multBy2(3))

	fmt.Println("applyUnaryOperation(3,createMultiplier(2)) =", applyUnaryOperation(3, createMultiplier(2)))

	// possible error handling
	q, r, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Divide(10, 0): Error %v\n", err)
	}
	fmt.Printf("Divide(10, 0): Quotient(%d), Remainder(%d)\n", q, r)

	doDivision(10, 0)
	doDivision(10, 3)
	msg, res := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(msg, res)

	msg, res = sum(1, 2, 3, 4, 5)
	fmt.Println(msg, res)

	slice := []int{1, 2, 3, 4, 5}
	msg, res = sum(slice...)
	fmt.Println("Variadic params as slice:", msg, res)
	deferred(1)

	fmt.Println("------------------ Panicking sample:")
	panicking(1)
	panicking(-10)
	fmt.Println("continuing execution after handled panicking...")

	os.Exit(0) // terminate program in particular point and provide exit code. 0 = ok
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func applyUnaryOperation(x int, operation func(int) int) int {
	return operation(x)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int { return x * factor }
}

func divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by 0!")
	}
	return a / b, a % b, nil
}

func divide2(a, b int) (quotient int, remainder int, err error) {
	if b == 0 {
		err = errors.New("Cannot divide by 0!")
	} else {
		quotient = a / b
		remainder = a % b
	}

	return
}

func doDivision(a, b int) (quotient int, remainder int) {
	var err error
	quotient, remainder, err = divide2(a, b)
	if err != nil {
		fmt.Printf("divide(%d, %d): Error - %v\n", a, b, err)
	} else {
		fmt.Printf("divide(%d, %d): quotient=%v, remainder=%v\n", a, b, quotient, remainder)
	}
	return quotient, remainder
}

func sum(nums ...int) (str string, total int64) {
	total = 0
	str = "THe sum of "
	for i, v := range nums {
		total = total + int64(v)
		if i > 0 {
			str += "+"
		}
		str += strconv.Itoa(v)
	}
	str += " ="
	return
}

func deferred(i int) {
	//NOTE:
	// 1. executed as FILO ordered
	// 2. Parameters evaluated at the moment of DEFINITION of the deferred statement

	defer fmt.Println("deferred print #1", i)

	i = 10

	defer fmt.Println("deferred print #2", i)

	i = 100
	defer fmt.Println("deferred print #3", i)

	i = 1000
	fmt.Println("Function execution sequence begin. i =", i)
	fmt.Println("Function execution sequence end. Exiting...")
}

func panicking(i int) {
	fmt.Printf("panicking(%v) entered\n", i)
	defer func() {
		fmt.Println("\tdeferred action #1", i)
		if r := recover(); r != nil {
			fmt.Println("\t  Recovering:", r) // r holds panic message
			fmt.Println("\t  Recover necessary resources here...")
		}
	}() // <-- NOTE! it is a function call!

	i++
	defer fmt.Println("\tdeferred action #2", i)
	i++
	fmt.Println("\tFunction execution sequence begin. i =", i)
	if i < 0 {
		fmt.Println("\tBefore panic(...)!")
		msg := "\t i < 0 (" + strconv.Itoa(i) + ") Something terrible happened..."
		panic(msg)
		//NOTE: no regular code after panic line will be called!
		//only defer code will be called before panicking... Use recover in deferred code
	}
	fmt.Println("\tYou should see this if i>0 at the moment of check!")
}
