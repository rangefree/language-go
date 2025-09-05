package main

import (
	"fmt"
	"log"
	"time"
)

func someOperation() {
	for i := range 20 {
		fmt.Println("iteration", i+1)
		time.Sleep(1 * time.Second)
	}
}

func tickerSample(maxVal int) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // must stop ticker to avoid resource leack
	i := 1

	// elegant way:
	// for range ticker.C {
	// 	i *= 2
	// 	fmt.Println("i =", i, "ticker:", ticker)
	// 	if i > maxVal {
	// 		break
	// 	}
	// }

	// Smart ass way:
	for {
		select {
		case <-ticker.C:
			i *= 2
			fmt.Println("i =", i, "ticker:", ticker)
			if i >= maxVal {
				fmt.Println("Exiting function.")
				return
			}

		}
	}
}

func main() {
	log.Println("main: entered")

	// timer := time.NewTimer(2 * time.Second) //non blocking!

	// v, ok := <-timer.C //blocking call
	// if !ok {
	// 	log.Println("main: timer expired. ")
	// } else {
	// 	log.Println("timer signaled.", v)
	// }

	timeout := time.After(2 * time.Second)
	//defer close(timeout) // no need to close read only channel
	done := make(chan bool)

	go func() {
		someOperation()
		done <- true
		close(done)
	}()

	//for {
	select {
	case <-timeout:
		fmt.Println("Timeout expired")

	case <-done:
		fmt.Println("operation complete")

	}
	//	}

	tickerSample(2048)

	log.Println("main exiting...")
}
