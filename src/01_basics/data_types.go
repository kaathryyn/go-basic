package main

import "fmt"

func main() {
	// Numeric types
	var integer int = 42
	var float float64 = 3.14
	var unsigned uint = 123
	
	// String type
	var text string = "Hello, Go!"
	
	// Boolean type
	var isGopher bool = true
	
	// Multiple declaration
	var (
		name    string = "Gopher"
		age     int    = 25
		isHappy bool   = true
	)
	
	// Constants
	const Pi = 3.14159
	const (
		StatusOK    = 200
		StatusError = 500
	)
	
	// Print all variables
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float)
	fmt.Printf("Unsigned: %d\n", unsigned)
	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Is Gopher: %t\n", isGopher)
	fmt.Printf("Name: %s, Age: %d, Is Happy: %t\n", name, age, isHappy)
	fmt.Printf("Pi: %.5f\n", Pi)
	fmt.Printf("Status OK: %d, Status Error: %d\n", StatusOK, StatusError)
	
	// Type conversion
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	
	fmt.Printf("\nType conversion:\n")
	fmt.Printf("int(%d) -> float64(%.1f) -> uint(%d)\n", i, f, u)
}