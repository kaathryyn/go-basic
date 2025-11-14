package main

import "fmt"

func main() {
	// Basic print statement
	fmt.Println("Hello, Go!")

	// Variables
	var name string = "Gopher"
	age := 25 // Type inference

	// Print with variables
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// Basic arithmetic
	x := 10
	y := 5
	sum := x + y
	difference := x - y
	product := x * y
	quotient := x / y

	// Print results
	fmt.Printf("Math with %d and %d:\n", x, y)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
}