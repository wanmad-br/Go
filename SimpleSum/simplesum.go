package main

import (
	"fmt"
)

// add is a function that takes two integers and returns their sum to user
func add(a int, b int) int {
	return a + b
}

func main() {
	// Declare variables
	num1 := 10 // Short variable declaration
	num2 := 5

	// Call the function
	result := add(num1, num2)

	// Print results to the console
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, result)
}
