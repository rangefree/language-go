package main

import (
	"fmt"
	"log"
	"time"
)

func sendMsg(sleepMs int, ch chan<- int, value int) {
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	log.Println("Seending value", value, "to the channel", ch)
	ch <- value
}

func channelSample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendMsg(90, ch1, 10)
	go sendMsg(20, ch2, 20)
	go sendMsg(1000, ch1, 11)
	go sendMsg(1070, ch2, 21)

	receiveduringSec := 3
	fmt.Println("channelSample: Begin receiving messages for", receiveduringSec, "second(s)...")
	startTime := time.Now()
	endTime := startTime.Add(time.Duration(receiveduringSec) * time.Second)

	for startTime.Before(endTime) {
		select {
		case msg := <-ch1:
			fmt.Println("channelSample: Receiving from channel1:", msg)
		case msg := <-ch2:
			fmt.Println("channelSample: Receiving from channel2:", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("channelSample: WARNING - Timeout signaled...")
			break
			//NOTE: timeout will not work if default: present!
			//default:
			//fmt.Println("channelSample: No channel ready...")
			//time.Sleep(100 * time.Millisecond)
		}
		startTime = time.Now()
	}
}

func main() {
	channelSample()
	fmt.Println()

	//--------------------------------------------
	selectInLoop()
	fmt.Println()

	//--------------------------------------------
	nonBlockingReceivwe()
	fmt.Println()

	//--------------------------------------------
	channelUsage()
	fmt.Println()

	//--------------------------------------------
	pipelineSample()
	fmt.Println()

	fmt.Println("main: exiting...")
}

func selectInLoop() {
	log.Println("selectInLoop: enterned")
	ch := make(chan int)
	go func() {
		log.Println("Will send 1")
		ch <- 1

		log.Println("Will send 2")
		ch <- 2

		log.Println("Will send 3")
		ch <- 3
		log.Println("Will close channel")
		close(ch)
		log.Println("goroutine: done")
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("selectInLoop: Channel is closed")
				// cleanup activities
				log.Println("selectInLoop: exiting")
				return
			}
			fmt.Println("selectInLoop: Received:", msg)
		}
	}

	log.Println("you should not see this.")
}

func nonBlockingReceivwe() {
	ch := make(chan int)

	select {
	case ch <- 1:
		fmt.Println("nonBlockingReceivwe: Sent to channel 1")
	default:
		fmt.Println("nonBlockingReceivwe: channel is not ready")
	}

	data := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case d := <-data:
				log.Println("nonBlockingReceivwe: got data:", d)
			case q := <-quit:
				log.Println("nonBlockingReceivwe: exiting:", q)
				if q {
					return
				}
			default:
				log.Println("nonBlockingReceivwe: waiting for data...")
				time.Sleep(500 * time.Millisecond)

			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(300 * time.Millisecond)
	}
	quit <- false
	quit <- true
	//close(data)
	//close(quit)
	time.Sleep(time.Second)
}

func channelUsage() {
	log.Println("channelUsage: entered")
	ch := make(chan int)

	go func() {
		for i := range 5 {
			log.Println("channelUsage: will write", i, "to channel")
			ch <- i
		}
		log.Println("channelUsage: closing channel")

		//NOTE: close channel on the writer side! if you try to close the closed channel then PANIC will happen
		close(ch) //no more writing to the channel
	}()

	for val := range ch {
		log.Println("channelUsage: received", val)
	}

	//same as above but uglier... ;-)
	for val, ok := <-ch; ok; val, ok = <-ch {
		log.Println("channelUsage: ERROR - reveiving valid value ", val, "on closed channel")
	}

	log.Println("channelUsage: done receiving data")
}

func pipelineSample() {
	log.Println("channelFilterring: entered")
	inCh := make(chan int)
	filteredCh := make(chan int)

	go producer(inCh, 5)
	go filter(inCh, filteredCh)

	for val := range filteredCh {
		log.Println("channelFilterring: finally received", val)
	}
	log.Println("channelFilterring: done")
}

func producer(ch chan<- int, count int) {
	log.Println("producer: entered")
	for i := range count {
		log.Println("producer: writing", i)
		ch <- i
		time.Sleep(time.Duration(i*200) * time.Millisecond)
	}
	log.Println("producer: closing channel")
	close(ch)
	log.Println("producer: done")
}

func filter(in <-chan int, out chan<- int) {
	log.Println("filter: entered")
	for val := range in {
		if val%2 == 0 {
			log.Println("filter: writing value", val)
			out <- val
		} else {
			log.Println("filter: filtering value", val, "out")
		}
	}
	log.Println("filter: closing filtered channel")
	close(out)
	log.Println("filter: done")
}
