package main

import (
	"fmt"
	"math"
)

// Shape interface defines a common behavior for all shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct represents a rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Methods for Rectangle implementing Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct represents a circle shape
type Circle struct {
	Radius float64
}

// Methods for Circle implementing Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Person struct demonstrates embedded structs and tags
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address   Address
}

// Address struct demonstrates nested structs
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Method for Person struct
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Constructor function for Person
func NewPerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// Employee demonstrates struct embedding
type Employee struct {
	Person   // Embedded struct
	JobTitle string
	Salary   float64
}

func main() {
	// Creating and using structs
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	// Using interface
	shapes := []Shape{rect, circle}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n",
			shape.Area(), shape.Perimeter())
	}

	// Creating and using nested structs
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Example City",
			Country: "Example Country",
		},
	}

	fmt.Printf("\nPerson: %s\n", person.FullName())
	fmt.Printf("Address: %s, %s\n", person.Address.Street, person.Address.City)

	// Using constructor function
	newPerson := NewPerson("Jane", "Smith", 25)
	fmt.Printf("New Person: %s\n", newPerson.FullName())

	// Using embedded structs
	emp := Employee{
		Person:   newPerson,
		JobTitle: "Software Engineer",
		Salary:   75000,
	}

	fmt.Printf("\nEmployee: %s\n", emp.FullName()) // Inherited from Person
	fmt.Printf("Job Title: %s\n", emp.JobTitle)
	fmt.Printf("Salary: $%.2f\n", emp.Salary)
}