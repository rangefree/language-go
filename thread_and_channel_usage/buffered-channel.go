package main

import (
	"fmt"
	"time"
)

func main() {
	// === blocking on send:
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Channel blocked (2 values)")
	//fmt.Println("Value: ", <-ch)
	//fmt.Println("Value: ", <-ch)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)

	}()

	ch <- 3 // if buffer is full => channel is blocked!
	fmt.Println("Channel unblocked")
	fmt.Println("Value:", <-ch)
	fmt.Println("Value:", <-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("Done...")

	// === Blocking on receive:

	ch1 := make(chan int, 2)
	go func() {
		fmt.Println("channel has no data")
		time.Sleep(2 * time.Second)
		fmt.Println("channel will have 1 now")
		ch1 <- 1

		time.Sleep(1 * time.Second)
		fmt.Println("channel will have 2 now")
		ch1 <- 2

	}()
	fmt.Println("Received:", <-ch1)
	fmt.Println("Received:", <-ch1)

	fmt.Println("Done...")

}
