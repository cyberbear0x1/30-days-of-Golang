// Basic Syntax, Data Types
// Challenge
// Declare variables of different data types (int, float64, string, bool), perform basic arithmetic operations, and print the results.

// Solution
package main

import "fmt"

func main() {
	// Declare variables of different data types
	var intVar int = 10
	var floatVar float64 = 3.5
	var stringVar string = "Hello, Go!"
	var boolVar bool = true

	// Perform basic arithmetic operations
	sum := intVar + int(5) // Adding 5 to intVar
	product := floatVar * 2.0
	division := floatVar / 2.0
	difference := intVar - 3

	// Print the variables and results
	fmt.Println("Integer Value:", intVar)
	fmt.Println("Float Value:", floatVar)
	fmt.Println("String Value:", stringVar)
	fmt.Println("Boolean Value:", boolVar)

	fmt.Println("Sum (intVar + 5):", sum)
	fmt.Println("Product (floatVar * 2.0):", product)
	fmt.Println("Division (floatVar / 2.0):", division)
	fmt.Println("Difference (intVar - 3):", difference)
}
