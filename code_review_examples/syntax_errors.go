package main

import (
	"fmt"
	"time
)

func main() {
	// Missing semicolon in declaration
	var x int y := 10
	
	// Mismatched types
	var name string = 42
	
	// Unclosed string
	fmt.Println("Hello, World)
	
	// Missing curly brace for if statement
	if x > 5
		fmt.Println("x is greater than 5")
	
	// Using := for already declared variable
	x := 20
	
	// Missing return value
	result := addNumbers(5, 10)
	fmt.Println(result)
}

// Function missing return statement despite return type
func addNumbers(a, b int) int {
	sum := a + b
	// Missing return statement
}
