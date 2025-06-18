package main

import "fmt"

func main() {
	var arr [5]int // initialized to 0!
	fmt.Println(arr)

	arr[0] = 100
	arr[4] = 200
	fmt.Println(arr)

	//var arr [5]string;
	letters := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(letters)
	orig := [3]int{1, 2, 3}
	copied := orig
	fmt.Println("orig =", orig)
	fmt.Println("copied =", copied)
	fmt.Println("are copied and source arrays equal ?", orig == copied)

	fmt.Println("Modify copied.")
	copied[1] = 100
	fmt.Println(orig)
	fmt.Println(copied)
	fmt.Printf("Length of the copied array is %d\n", len(copied))

	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[", i, "] = ", arr[i])
	}

	for i, v := range arr {
		fmt.Printf("value @ index %d = %d\n", i, v)
	}

	//blank identifier useful when returned value should be ignored
	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}

	b := 1
	_ = b //GO specific craziness :)

	//multidimentional arrays:
	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("matrix =", matrix)
	fmt.Println("len(matrix) =", len(matrix), "!")

	for i, v := range matrix {
		fmt.Printf("value @ index %d = %d\n", i, v)
	}
}
