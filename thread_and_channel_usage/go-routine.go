package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("HELLO from main thread!")

	go sayHello() // execute function in separate thread (uses M:N scheduling model)
	// return   // sample of goroutine leak

	fmt.Println("main thread is doing something for 2 sec.")

	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("messge to log")
	go printNumbers()
	go printLetters()

	var err error

	// the way to call goroutines which return errors
	go func() {
		err = couldThrowError()
	}()

	time.Sleep(6 * time.Second)

	// challenge is to check error AFTEER gorutine is finished (!) - not tri8vial in real life.
	if err != nil {
		fmt.Println("Error happened:", err)
	}

	fmt.Println("main thread is Exiting")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("HELLO from goroutine!")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(time.Now(), i)
		time.Sleep((500 * time.Millisecond))
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(time.Now(), string(letter))
		time.Sleep((1 * time.Second))
	}
}

func couldThrowError() error {
	time.Sleep((1 * time.Second))
	return fmt.Errorf("XEP BAM!")
}
