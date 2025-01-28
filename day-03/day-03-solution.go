// Challenge
// Control Flow (if/else, for loops, switch)
// Write a program that checks if a given number is even or odd.

package main

import "fmt"

func main() {
	Odd_or_Even(4)
}

// Create a function to check if a given number is even or odd.
func Odd_or_Even(x int) {
	if x%2 == 0 {
		fmt.Println(x, "is even.")
	} else {
		fmt.Println(x, "is odd.")
	}
}
