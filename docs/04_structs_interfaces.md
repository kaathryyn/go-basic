# Structs and Interfaces in Go

## Structs

Structs are user-defined types that group together related data fields. They are the foundation of object-oriented programming in Go.

### Basic Struct Definition
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```

### Creating Struct Instances
```go
// Method 1: Declaring all fields
person1 := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// Method 2: Short declaration (order matters)
person2 := Person{"Jane", "Smith", 25}

// Method 3: Zero-valued struct
var person3 Person
```

### Struct Tags
Struct tags provide metadata about fields:
```go
type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}
```

### Nested Structs
```go
type Address struct {
    Street  string
    City    string
    Country string
}

type Person struct {
    Name    string
    Address Address
}
```

### Embedded Structs (Composition)
```go
type Employee struct {
    Person   // Embedded struct
    JobTitle string
    Salary   float64
}
```

## Interfaces

Interfaces define behavior through method signatures. They enable polymorphism in Go.

### Interface Definition
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Implementing Interfaces
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

### Empty Interface
```go
interface{} // Can hold values of any type
```

## Best Practices

### 1. Struct Design
- Keep structs focused and cohesive
- Use embedding for composition over inheritance
- Make fields public (capitalized) only when necessary
- Use constructor functions for complex initialization

### 2. Interface Design
- Keep interfaces small and focused
- Follow the interface segregation principle
- Define interfaces where they're used
- Use composition of interfaces for complex behaviors

### 3. Method Receivers
- Use pointer receivers when:
  - Methods need to modify the receiver
  - Struct is large
  - Consistency is required with other methods
- Use value receivers when:
  - Methods don't modify the receiver
  - Struct is small
  - Immutability is desired

### 4. Error Handling
- Use custom error types for specific error cases
- Implement the error interface for custom error types
- Return errors as the last return value

## Common Patterns

### 1. Constructor Pattern
```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

### 2. Builder Pattern
```go
type PersonBuilder struct {
    person *Person
}

func (b *PersonBuilder) WithName(name string) *PersonBuilder {
    b.person.Name = name
    return b
}
```

### 3. Interface Composition
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Testing Structs and Interfaces

### 1. Table-Driven Tests
```go
func TestArea(t *testing.T) {
    tests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{Width: 4, Height: 5}, 20},
        {Circle{Radius: 3}, 28.27},
    }

    for _, tt := range tests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
    }
}
```

### 2. Interface Mocking
```go
type DataStore interface {
    Save(data []byte) error
}

type MockDataStore struct {
    SaveFunc func(data []byte) error
}

func (m *MockDataStore) Save(data []byte) error {
    return m.SaveFunc(data)
}
```