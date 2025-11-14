package main

import "fmt"

func main() {
	// If-else statement
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with a short statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or lower")
	}

	// For loop - basic
	fmt.Println("\nCounting to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop as while
	fmt.Println("\nWhile-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// Range-based for loop
	fmt.Println("\nIterating over slice:")
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// Switch statement
	fmt.Println("\nSwitch example:")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek")
	}

	// Switch with no condition (like if-else)
	fmt.Println("\nGrade evaluation:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 80:
		fmt.Println("Good")
	case score >= 70:
		fmt.Println("Fair")
	default:
		fmt.Println("Needs improvement")
	}
}