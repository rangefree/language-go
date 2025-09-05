package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	fmt.Println("Worker", id, "started")
	for task := range tasks {
		fmt.Printf("Worker %d: executing task %d\n", id, task)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d, task %d: returning result.\n", id, task)
		results <- task * 2
	}
	fmt.Println("Worker", id, "exiting...")
}

func main() {
	numOfWorkers := 3
	numOfJobs := 10

	tasks := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	// Create pool of workers:
	for workerId := range numOfWorkers {
		go worker(workerId, tasks, results)
	}

	// sent tasks to the channel:
	for i := range numOfJobs {
		tasks <- i
	}
	close(tasks)

	for jobId := range numOfJobs {
		result := <-results
		fmt.Println("result from job", jobId, " is", result)
	}
	close(results)

}
