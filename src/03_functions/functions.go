package main

import (
	"fmt"
	"strings"
)

// Basic function
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Function with multiple parameters and return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Function with named return values
func getMinMax(numbers []int) (min, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return // naked return uses named return values
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as parameter
func processString(s string, processor func(string) string) string {
	return processor(s)
}

// Closure (anonymous function)
func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Basic function call
	message := greet("Gopher")
	fmt.Println(message)

	// Multiple return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// Named return values
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := getMinMax(numbers)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d\n", total)

	// Function as parameter
	upperCase := processString("hello", strings.ToUpper)
	fmt.Println("Processed string:", upperCase)

	// Closure
	counter := getCounter()
	fmt.Println("Count:", counter()) // 1
	fmt.Println("Count:", counter()) // 2
	fmt.Println("Count:", counter()) // 3
}