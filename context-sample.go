package main

import (
	"context"
	"log"
	"time"
)

func doWork(ctx_ context.Context) {
	log.Println("doWork: entered")
	for {
		select {
		case <-ctx_.Done():
			log.Println("doWork: was canceled for", ctx_.Value("id"), ". Reason:", ctx_.Err())
			return

		default:
			log.Println("doWork: working for", ctx_.Value("id"), "...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx, doneOnTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer doneOnTimeout() // this function will be called at the moment when timeout will be reached...
	ctx = context.WithValue(ctx, "id", "John")

	go doWork(ctx)
	log.Println("Begin waiting...")
	time.Sleep(3 * time.Second)
	log.Println("Done waiting...")

	ctx1, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel() //can be called manually to signal cancelation in the done state
	}()

	go doWork(ctx1)
	time.Sleep(3 * time.Second)

}
