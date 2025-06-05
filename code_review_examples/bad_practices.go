package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Global variables - generally discouraged in Go
var GlobalCounter int = 0
var CONNECTION_STRING string = "mongodb://user:password@localhost:27017" // Hard-coded credentials

func main() {
	// Ignoring errors - bad practice
	file, _ := os.Open("file.txt")
	defer file.Close()
	
	data, _ := ioutil.ReadAll(file)
	fmt.Println(string(data))
	
	// Magic numbers
	if len(data) > 1024 {
		fmt.Println("File is larger than expected")
	}
	
	// Inefficient string concatenation in a loop
	var result string
	for i := 0; i < 1000; i++ {
		result = result + "a" // Should use strings.Builder instead
	}
	
	// Not checking for nil
	var ptr *string = nil
	processString(ptr)
	
	// Unnecessary else after return
	value := getValue()
	if value > 100 {
		fmt.Println("Large value")
		return
	} else {
		fmt.Println("Small value")
	}
}

// Unexported function with exported parameters - inconsistent naming
func processString(Input *string) {
	fmt.Println(*Input) // Will panic if nil
}

// Function that could return an error but doesn't
func getValue() int {
	// Should handle potential errors
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	return len(body)
}
