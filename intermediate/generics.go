package main

import "fmt"

func swap[T any](a, b *T) {
	if a == b {
		return
	}
	temp := *a
	*a = *b
	*b = temp
}

type Stack[T any] struct {
	elements []T
}

func (a *Stack[T]) Push(element T) {
	a.elements = append(a.elements, element)
}

func (a *Stack[T]) Pop() (T, bool) {
	pos := len(a.elements) - 1 // last index
	if pos < 0 {
		var NIL T
		return NIL, false
	}

	ret := a.elements[pos]        // copy last
	a.elements = a.elements[:pos] // deleting last element
	return ret, true
}

func (a Stack[T]) IsEmpty() bool {
	return len(a.elements) == 0
}

func PRINT[T any](v T) {
	fmt.Println(v)
}

func main() {
	v1 := "aaa"
	v2 := "bbb"

	println(v1, v2)
	swap(&v1, &v2)
	println("After swap:", v1, v2)

	stackStr := Stack[string]{}
	stackStr.Push("one")
	stackStr.Push("two")
	stackStr.Push("three")
	fmt.Println(stackStr)
	tryPop := func() {
		v, ok := stackStr.Pop()
		if !ok {
			fmt.Printf("Failed to Pop() element from the stack\n")
		} else {
			fmt.Printf("Popped element (%v), Stack condition: %v\n", v, stackStr)
		}
	}

	tryPop()
	tryPop()
	tryPop()
	tryPop()
	if stackStr.IsEmpty() {
		fmt.Println("Stack is Empty!")
	}

	// v, ok := stackStr.Pop()
	// if !ok {
	// 	fmt.Println("Failed to Pop() element from the stack")
	// } else {
	// 	fmt.Println("Popped element (%v), Stack condition: %v\n", v, stackStr)
	// }

	PRINT([]int{1, 2, 3})

}
