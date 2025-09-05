package main

import (
	"fmt"
	"time"
)

func main() {
	greetingChannel := make(chan string) // create unbuffered channel (requires immediate receiver!)
	greetString := "Hello from Channel"

	//NOTE: go channels are Blocking (!)
	// channels must be used inside goroutines
	go func() {
		greetingChannel <- greetString // sending value to the channel
		greetingChannel <- "Whatever! (second receive required...)"
	}()

	go func() {
		received := <-greetingChannel // receiving value from the channel (also blocking!)
		fmt.Println(received)

		received = <-greetingChannel // secong receiving
		fmt.Println(received)
	}()

	go func() {
		for _, e := range "ABCDE" {
			greetingChannel <- "Character " + string(e)
		}
	}()

	for _ = range 5 {
		str := <-greetingChannel
		fmt.Println(str)
	}

	ch := make(chan int)
	go func() {
		ch <- 1 // need sceiver => we send data in goroutine!
	}()
	receiver := <-ch //receier should be in different goroutine than sender for unbuffered channel
	fmt.Println(receiver)
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
}
