# Go Programming Basics

## Introduction to Go

Go (or Golang) is a statically typed, compiled programming language designed by Google. It combines the efficiency of compiled languages with the ease of use of dynamic languages.

## Key Features of Go

- Simple and clean syntax
- Built-in concurrency support
- Strong standard library
- Fast compilation
- Garbage collection
- Built-in testing support

## Basic Syntax Elements

### Package Declaration
Every Go file starts with a package declaration:
```go
package main
```

### Import Statements
Import other packages using the `import` keyword:
```go
import "fmt"
```

### Main Function
The entry point of a Go program is the `main` function in the `main` package:
```go
func main() {
    // Your code here
}
```

### Variables

Go supports multiple ways to declare variables:

1. Using `var`:
```go
var name string = "Gopher"
var age int = 25
```

2. Short declaration (type inference):
```go
name := "Gopher"
age := 25
```

### Basic Data Types

- **Numeric Types:**
  - int, int8, int16, int32, int64
  - uint, uint8, uint16, uint32, uint64
  - float32, float64
  - complex64, complex128
- **String Type:** string
- **Boolean Type:** bool

### Printing Output

The `fmt` package provides various printing functions:

```go
fmt.Println("Hello") // Prints with newline
fmt.Print("Hello")   // Prints without newline
fmt.Printf("Hello, %s", name) // Formatted print
```

## Running Go Programs

1. Navigate to your code directory
2. Run with: `go run filename.go`
3. Build executable: `go build filename.go`

## Code Style

- Go enforces a strict code format
- Use `go fmt` to automatically format your code
- Follow the official Go style guide
- Use meaningful variable names
- Add comments for clarity

## Next Steps

1. Try modifying the example code
2. Experiment with different data types
3. Practice writing simple programs
4. Move on to control structures (if/else, loops)