package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (pS *Counter) incr() {
	pS.m.Lock()
	defer pS.m.Unlock()
	pS.value++
}

func (pS *Counter) decr() {
	pS.m.Lock()
	defer pS.m.Unlock()
	pS.value--
}

func (pS *Counter) get() int {
	pS.m.Lock()
	defer pS.m.Unlock()
	return pS.value
}

type Progress struct {
	mtx     sync.Mutex
	final   int
	current int
	size    uint8
	done    bool
}

func (pS *Progress) init(maxValue int, size uint8) {
	pS.mtx.Lock()
	defer pS.mtx.Unlock()
	pS.final = maxValue
	pS.current = 0
	pS.size = size
	pS.done = false
}

func (pS *Progress) show() {
	pS.mtx.Lock()
	defer pS.mtx.Unlock()

	if pS.size == 0 || pS.done {
		return
	}

	//final   - size
	//current - doneCount
	//doneCount := current * size / final

	var doneCount uint8
	doneCount = uint8(pS.current * int(pS.size) / pS.final)

	if pS.final >= pS.current {
		//doneCount = uint8(pS.current * int(pS.size) / pS.final)
		p := float32(float64(pS.current*100) / float64(pS.final))

		fmt.Print("\rProgres:")
		for i := range pS.size {
			if i > doneCount {
				fmt.Print("-")
			} else {
				fmt.Print("X")
			}
		}

		if pS.final > pS.current {
			fmt.Printf("(%3.2f%%)", p)
		} else {
			fmt.Println("(100%)")
		}
	}

}

func (pS *Progress) add(value int) {
	pS.mtx.Lock()
	defer pS.mtx.Unlock()

	if pS.current < pS.final {
		pS.current += value
	}
}

// var wheelMtx sync.Mutex
// var wheelState int = 0

// func moveWheel() {
// 	wheelMtx.Lock()
// 	defer wheelMtx.Unlock()
// 	fmt.Print("\r")
// 	switch wheelState {
// 	case 0:
// 		fmt.Print(" | ")
// 	case 1:
// 		fmt.Print(" x ")
// 	case 2:
// 		fmt.Print(" - ")
// 	case 3:
// 		fmt.Print(" x ")
// 	case 4:
// 		fmt.Print(" + ")
// 	}
// 	wheelState++
// 	time.Sleep(time.Second)
// 	if wheelState > 4 {
// 		wheelState = 0
// 	}

// }

func sampleConcurentIncr(numWorkers int, numRepsPerWorker int) {
	var wg sync.WaitGroup
	counter := &Counter{}
	wg.Add(numWorkers)
	var progress Progress
	progress.init(numWorkers*numRepsPerWorker, 10)
	for id := range numWorkers {
		go func() {
			fmt.Println("worker", id, "entered")
			defer wg.Done()
			for i := range numRepsPerWorker {
				counter.incr()
				progress.add(1)
				//progress.show()
				time.Sleep(10 * time.Millisecond)
				if i%10 == 0 {
					progress.show()
					//moveWheel()
				} //fmt.Printf(".")
			}
			progress.show()
			fmt.Println("worker", id, "exiting...")
		}()
	}

	wg.Wait()

	fmt.Println("Final Count value is", counter.value)
	if counter.value != numRepsPerWorker*numWorkers {
		fmt.Println("Something went wrong. Expected value is", numRepsPerWorker*numWorkers)
	}
}

// ---------------------------------------------------
type AtomicCounter struct {
	value int64
}

func (pS *AtomicCounter) add(i int64) {
	atomic.AddInt64(&pS.value, i)
}

func (pS *AtomicCounter) get() int64 {
	return atomic.LoadInt64(&pS.value)
}

func sampleAtomicCounter(numWorkers int, numRepsPerWorker int) {
	var wg sync.WaitGroup
	counter := &AtomicCounter{}
	wg.Add(numWorkers)
	var progress Progress
	progress.init(numWorkers*numRepsPerWorker, 20)
	for id := range numWorkers {
		go func() {
			fmt.Println("worker", id, "entered")
			defer wg.Done()
			for i := range numRepsPerWorker {
				counter.add(1)
				progress.add(1)
				time.Sleep(10 * time.Millisecond)
				if i%10 == 0 || i == numRepsPerWorker-1 {
					progress.show()
				}
			}
			fmt.Println("worker", id, "exiting...")
		}()
	}

	wg.Wait()

	fmt.Println("Final Count value is", counter.get()) // it is ok to access counter.value directly here because no threrads are active at this time
	if int(counter.get()) != numRepsPerWorker*numWorkers {
		fmt.Println("Something went wrong. Expected value is", numRepsPerWorker*numWorkers)
	}
}

func main() {
	numWorkers := 10
	numRepsPerWorker := 1000

	fmt.Println("Sample of the manual syncronization:")
	sampleConcurentIncr(numWorkers, numRepsPerWorker)

	fmt.Println("Sample of the atomic syncronization:")
	sampleAtomicCounter(numWorkers, numRepsPerWorker)

	fmt.Println("main(): exiting...")
}
