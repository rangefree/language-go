package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, pWaitGroup *sync.WaitGroup, res chan<- int) {
	defer pWaitGroup.Done()
	fmt.Println("Worker", id, "starting")

	for i := range 5 {
		i++
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Worker", id, "running for", i*500, "ms")
	}
	res <- id * 1000

	fmt.Println("Worker", id, "exiting...")
}

func main() {
	fmt.Println("main(): starting")
	numWorkers := 4
	var wg sync.WaitGroup

	results := make(chan int, numWorkers)
	wg.Add(numWorkers) // do it Before workers are started! :-)
	for i := range numWorkers {
		go worker(i, &wg, results)
	}

	go func() {
		fmt.Println("...(): begin waiting")
		wg.Wait() // blocking
		fmt.Println("...(): done waiting")
		close(results)
		fmt.Println("...(): results channel is closed")

	}()

	for res := range results {
		fmt.Println("main(): worker's result is", res)
	}
	fmt.Println("main(): exiting...")
}
