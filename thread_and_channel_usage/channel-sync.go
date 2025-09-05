package thread_and_channel_usage

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main: entered")
	done := make(chan bool)

	go func() {
		fmt.Println("enter routine")
		time.Sleep(1 * time.Second)
		done <- true
	}()

	// next receiving will block main thread until goroutine is done:
	if <-done {
		fmt.Println("main: Rooutin finished successfully.")
	}
	fmt.Println()

	//---------------------------------------------------------------------
	ch1 := make(chan int)
	go func() {
		fmt.Println("enter routine")
		fmt.Println("routine: will send value")
		ch1 <- 9
	}()

	// next statement will block until gorotine will send a value:
	fmt.Println("received:", <-ch1)
	fmt.Println()

	//---------------------------------------------------------------------
	// emmulating wait for all started goroutines
	numRoutines := 3
	done1 := make(chan int, numRoutines)
	for i := range numRoutines {
		go func(id int) {
			fmt.Println("Routine", id, "started")
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Println("Routine", id, "exiting")
			done1 <- id
		}(i + 1)
	}

	for _ = range numRoutines {
		fmt.Println("Routine", <-done1, "signaled finish")
	}
	fmt.Println()

	//---------------------------------------------------------------------
	//auto generated receivers
	data := make(chan string) //bi-directional
	// go func(ch chan<- string) { //write only channel (in)
	// 	for i := range 5 {
	// 		data <- "data block " + string(i)
	// 		time.Sleep((500 * time.Millisecond))
	// 	}
	// 	close(data) //will let GO know safe size of the range(data)
	// }(data)

	// for value := range data { // automatic way of creating receiver each iteration
	// 	fmt.Println("Received:", value, time.Now())
	// }
	go producer(data)
	consumer(data)

	fmt.Println("main: exiting...")
}

func producer(ch chan<- string) {
	for i := range 5 {
		ch <- "data block " + string('0'+i)
		time.Sleep((500 * time.Millisecond))
	}
	close(ch) //will let GO know safe size of the range(data)
}

func consumer(ch <-chan string) {
	for value := range ch { // automatic way of creating receiver each iteration
		fmt.Println("Received:", value, time.Now())
	}
}
