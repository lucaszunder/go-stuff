package main

import (
	"fmt"
	"sync"
)

func main() {
	answers := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go add(1, 2, &wg, answers)
	go add(3, 2, &wg, answers)
	go add(4, 4, &wg, answers)

	go func() {
		wg.Wait()
		close(answers)
	}()

	for value := range answers {
		fmt.Println(value)
	}
}

func add(a, b int, wg *sync.WaitGroup, output chan int) {
	defer wg.Done()
	output <- a + b
}

// process M -> Thread | OS cycle

// go routine run inside a process, and for this process we will allocate one thread.
// multiple goroutines != multiple threads, we will only run in multiple threads if one the main thread is blocked by something

// scheduler is the responsible for change the thread state
// a time is generated for a process and split equally for every thread by the scheduler
// possible problems with multiple threads Context Switching - C10k problem - Fixed Stack size
// thread runnable - running - waiting
