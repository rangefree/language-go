package main

import "fmt"

func main() {
	var slice []int // it is std::vector<int>

	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)
	copySlice := slice
	fmt.Println("copied from slice:", copySlice)
	copySlice[0] = 100
	fmt.Println("modify copied slice:", copySlice)
	fmt.Println("original slice:", slice)

	slice1 := make([]int, 5)
	fmt.Println("auto generated slice:", slice1)
	a := [5]int{1, 2, 3, 4, 5}
	slice2 := a[1:4] // [1,4) !
	fmt.Println("source array:", a)
	fmt.Println("generated slice2 from array[1:4]:", slice2)
	slice2 = append(slice2, 1, 5)
	fmt.Println("slice2 = append(slice2, 1,5) = ", slice2)

	pSlice := &slice2
	fmt.Println("slice2 = ", slice2)
	fmt.Println("pSlice = ", pSlice)
	fmt.Println("*pSlice = ", *pSlice)

	var nilSlice []int // nil slice
	if nilSlice == nil {
		fmt.Println("We have nil slice")
	}

	// multidim slice:
	numRows := 4
	twoD := make([][]int, numRows)
	for i := range numRows {
		rowLen := i + 1
		twoD[i] = make([]int, rowLen)
		for j := range rowLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	//var slice2 string = "aaaaa"

	src := []int{1, 2, 3, 4, 5}
	dst := src[1:4]
	fmt.Println("src =", src)
	fmt.Println("dst =", dst)

	fmt.Println("modify dst[0]")
	dst[0] = 100
	fmt.Println("src =", src)
	fmt.Println("dst =", dst)
	fmt.Println("len(src) =", len(src))
	fmt.Println("len(dst) =", len(dst))
	//NOTE: Slice is a reference to the underlining array...
}
