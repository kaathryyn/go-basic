# Functions in Go

Functions are the building blocks of Go programs. They encapsulate reusable code and provide modularity to your programs.

## Function Declaration

Basic syntax for declaring a function:
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

## Function Types

### 1. Basic Functions
```go
func greet(name string) string {
    return "Hello, " + name
}
```

### 2. Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### 3. Named Return Values
```go
func getMinMax(numbers []int) (min, max int) {
    // variables min and max are declared automatically
    // return statement will return their current values
    return
}
```

### 4. Variadic Functions
Functions that accept variable number of arguments:
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## Advanced Function Concepts

### 1. Functions as Values
In Go, functions are first-class citizens:
```go
var compute func(int, int) int
compute = func(x, y int) int {
    return x + y
}
```

### 2. Closures
Anonymous functions that can access variables from the enclosing scope:
```go
func getCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 3. Methods
Functions associated with a type:
```go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

## Best Practices

1. **Function Names**
   - Use camelCase
   - Should be descriptive
   - Start with lowercase letter for package-level functions
   - Start with uppercase letter for exported functions

2. **Parameters**
   - Group parameters of the same type
   - Use meaningful parameter names
   - Consider using structs for many parameters

3. **Return Values**
   - Use named returns for clarity
   - Return errors as the last return value
   - Use multiple returns judiciously

4. **Error Handling**
   - Always handle returned errors
   - Return errors rather than panicking
   - Use error as the last return value

5. **Documentation**
   - Write godoc comments for exported functions
   - Start with the function name
   - Document parameters and return values

## Common Patterns

1. **Error Handling Pattern**
```go
result, err := someFunction()
if err != nil {
    return err
}
```

2. **Builder Pattern**
```go
func NewObject() *Object {
    return &Object{}
}

func (o *Object) WithField(field string) *Object {
    o.field = field
    return o
}
```

3. **Options Pattern**
```go
type Option func(*Config)

func WithTimeout(t time.Duration) Option {
    return func(c *Config) {
        c.timeout = t
    }
}
```

## Testing Functions

- Place tests in `*_test.go` files
- Use `testing` package
- Write table-driven tests
- Test both success and failure cases