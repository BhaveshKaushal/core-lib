package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Memory leak - resources not properly cleaned up
	leakyFunction()
	
	// Inefficient slice usage - pre-allocating would be better
	processLargeData()
	
	// Copying large structs instead of using pointers
	inefficientStructCopy()
	
	// Improper defer usage inside a loop
	for i := 0; i < 1000; i++ {
		processWithDefer(i)
		// Each iteration adds a defer call that only executes when the function returns
	}
}

func leakyFunction() {
	// Creating a channel but never consuming from it
	dataChan := make(chan []byte)
	
	// Goroutine that has no way to exit - memory leak
	go func() {
		bigData := make([]byte, 10*1024*1024) // 10MB of data
		for {
			dataChan <- bigData // This will block forever if no one reads
		}
	}()
}

func processLargeData() {
	// Inefficient way to build a large slice
	data := []int{}
	for i := 0; i < 100000; i++ {
		// Causes many reallocations and copies
		data = append(data, i)
	}
	
	// Better approach would be:
	// data := make([]int, 0, 100000)
}

func inefficientStructCopy() {
	type LargeStruct struct {
		Data     [1024]int
		Metadata map[string]string
		mu       sync.Mutex
	}
	
	original := LargeStruct{
		Metadata: make(map[string]string),
	}
	
	// Fill the struct
	for i := 0; i < 1024; i++ {
		original.Data[i] = i
		original.Metadata[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}
	
	// Inefficient: Copying the entire large struct
	copy := original
	
	// Modify copy
	copy.Data[0] = 999
	
	// Force GC for demonstration
	runtime.GC()
}

func processWithDefer(i int) {
	// Resource allocation
	resource := acquireExpensiveResource()
	
	// Defer inside a loop is inefficient - all defers wait until the function returns
	defer releaseExpensiveResource(resource)
	
	// Process the resource
	fmt.Println("Processing resource:", i)
	
	// Better approach: release immediately after using in a loop context
	// releaseExpensiveResource(resource)
}

func acquireExpensiveResource() string {
	return "expensive-resource"
}

func releaseExpensiveResource(resource string) {
	// Simulate resource cleanup
}
