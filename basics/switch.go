package main

import "fmt"

func main() {
	var i int
	CheckType(i)       //int
	CheckType("hello") //string
	CheckType(3.14)
	CheckType(true)
}

func CheckType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println(x, "is Integer")
		//fallthrough does not work in type switch !
	case int16:
		fmt.Println(x, "is int 16-bits")
	case int32:
		fmt.Println(x, "is int 32-bits")
	case float64:
		fmt.Println(x, "is Float")
	case string:
		fmt.Println(x, "is String")
	default:
		fmt.Println(x, "is Whatever it is...")
	}
}
