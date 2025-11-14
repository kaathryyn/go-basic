package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Custom error type
type ValidationError struct {
	Field string
	Error string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field %s: %s", e.Field, e.Error)
}

// Function that returns a basic error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function that returns a custom error
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field: "age",
			Error: "age cannot be negative",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field: "age",
			Error: "age seems unrealistic",
		}
	}
	return nil
}

// Function that wraps errors
func processAge(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("age validation failed: %w", err)
	}
	return nil
}

// Function demonstrating panic and recovery
func performDangerousOperation(value int) (err error) {
	// Defer a recovery function
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	if value < 0 {
		panic("negative value not allowed")
	}

	return nil
}

// Function demonstrating file operations with error handling
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Always close the file

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func main() {
	// Basic error handling
	result, err := divide(10, 0)
	if err != nil {
		log.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// Custom error handling
	ages := []int{25, -5, 200}
	for _, age := range ages {
		if err := processAge(age); err != nil {
			// Type assertion to check for specific error type
			var validationErr *ValidationError
			if errors.As(err, &validationErr) {
				log.Printf("Validation error: %v\n", validationErr)
			} else {
				log.Printf("Other error: %v\n", err)
			}
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}

	// Panic recovery
	err = performDangerousOperation(-1)
	if err != nil {
		log.Printf("Recovered error: %v\n", err)
	}

	// File operation error handling
	err = writeToFile("test.txt", "Hello, Go!")
	if err != nil {
		log.Printf("File operation error: %v\n", err)
	}

	// Error wrapping and unwrapping
	err = processAge(-5)
	if err != nil {
		// Get the underlying error
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Original error: %v\n", validationErr)
		}
		
		// Check if error is of a specific type
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		}
	}
}