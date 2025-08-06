package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(rand.Intn(10)) //auto seeded randome numbers

	seed := int64(42)
	gen := rand.New(rand.NewSource(seed))

	//Generate nums 1-10:
	fmt.Printf("Generating random numbers for seed %d: ", seed)
	for range 10 {
		fmt.Print(gen.Intn(10)+1, " ")
	}
	fmt.Println()

	seed1 := time.Now().Unix()
	gen1 := rand.New((rand.NewSource(seed1)))
	fmt.Printf("Generating random numbers for seed %d: ", seed1)
	for range 10 {
		fmt.Print(gen1.Intn(10)+1, " ")
	}
	fmt.Println()

	RollDices()
}

func RollDices() {
	for {
		fmt.Println()
		fmt.Println("1. Roll 2 dices:")
		fmt.Println("0. Exit")
		fmt.Println()

		var choice uint8
		_, err := fmt.Scan(&choice)

		fmt.Println()
		if err != nil || choice > 1 || choice == 0 {
			break
		}
		d1 := rand.Intn(6) + 1
		d2 := rand.Intn(6) + 1

		fmt.Printf("Dices: [%d] [%d] : %d\n", d1, d2, d1+d2)
		fmt.Println(strings.Replace("go gopher", "go", "Go", 1))
		layout := "2006-01-02 15:04:05"
		str := "2024-07-03 14:28:35"
		t, _ := time.Parse(layout, str)
		fmt.Println(t)
		
	}
}
