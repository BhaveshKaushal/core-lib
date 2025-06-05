package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared variable without proper synchronization
var counter int = 0

func main() {
	// Race condition - multiple goroutines accessing shared data without locks
	for i := 0; i < 1000; i++ {
		go incrementCounter() // Race condition
	}
	
	// Deadlock situation
	var mutex1, mutex2 sync.Mutex
	
	go func() {
		mutex1.Lock()
		time.Sleep(100 * time.Millisecond)
		mutex2.Lock() // Potential deadlock if another goroutine locks mutex2 first and then tries to lock mutex1
		// Critical section
		mutex2.Unlock()
		mutex1.Unlock()
	}()
	
	go func() {
		mutex2.Lock()
		time.Sleep(100 * time.Millisecond)
		mutex1.Lock() // Deadlock!
		// Critical section
		mutex1.Unlock()
		mutex2.Unlock()
	}()
	
	// Incorrect WaitGroup usage
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		go func(i int) {
			// Forgot to call wg.Add() before launching goroutine
			defer wg.Done() // This might cause a panic if wg counter is 0
			fmt.Println("Processing:", i)
		}(i)
	}
	wg.Wait()
	
	// Common channel mistake - sending on a nil channel will block forever
	var ch chan int // nil channel
	go func() {
		ch <- 1 // Will block forever
	}()
	
	time.Sleep(time.Second)
	fmt.Println("Final counter:", counter)
}

func incrementCounter() {
	temp := counter
	// Simulating some processing time that makes race condition more likely
	time.Sleep(time.Microsecond)
	counter = temp + 1
}
