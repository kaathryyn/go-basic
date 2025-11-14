# Error Handling in Go

Go takes a unique approach to error handling, treating errors as values that can be returned, checked, and passed around.

## Basic Error Handling

### The error Interface
```go
type error interface {
    Error() string
}
```

### Returning Errors
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Checking Errors
```go
result, err := divide(10, 0)
if err != nil {
    log.Printf("Error: %v\n", err)
    return
}
```

## Custom Error Types

### Creating Custom Errors
```go
type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

### Using Custom Errors
```go
if err := validate(); err != nil {
    var validErr *ValidationError
    if errors.As(err, &validErr) {
        // Handle validation error
    }
}
```

## Error Wrapping

### Wrapping Errors
```go
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

### Unwrapping Errors
```go
errors.Unwrap(err)
errors.Is(err, targetErr)
errors.As(err, &targetErr)
```

## Best Practices

1. **Error Types**
   - Use custom error types for specific error cases
   - Implement the error interface
   - Include relevant context in error messages

2. **Error Handling**
   - Check errors immediately
   - Don't ignore errors
   - Wrap errors to add context
   - Use type assertions when needed

3. **Error Messages**
   - Be descriptive but concise
   - Include relevant values
   - Use consistent formatting
   - Avoid redundant information

4. **Error Propagation**
   - Wrap errors when adding context
   - Don't log and return
   - Consider creating error packages for domains

## Common Patterns

### Sentinel Errors
```go
var (
    ErrNotFound = errors.New("not found")
    ErrInvalidInput = errors.New("invalid input")
)
```

### Type-Based Error Checking
```go
type NotFoundError struct {
    Item string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Item)
}
```

### Multiple Error Handling
```go
type MultiError struct {
    Errors []error
}

func (m *MultiError) Error() string {
    var errStrings []string
    for _, err := range m.Errors {
        errStrings = append(errStrings, err.Error())
    }
    return strings.Join(errStrings, "; ")
}
```

## Panic and Recovery

### When to Panic
- Unrecoverable errors
- Programming errors
- Initialization failures

### Panic Handler
```go
defer func() {
    if r := recover(); r != nil {
        log.Printf("Recovered from panic: %v", r)
    }
}()
```

## Error Handling in Tests

### Testing for Errors
```go
func TestDivide(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Error("expected error but got none")
    }
    if !errors.Is(err, ErrDivideByZero) {
        t.Errorf("expected ErrDivideByZero but got %v", err)
    }
}
```

### Table-Driven Error Tests
```go
tests := []struct {
    name    string
    input   string
    wantErr bool
}{
    {"valid input", "valid", false},
    {"invalid input", "", true},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        err := validate(tt.input)
        if (err != nil) != tt.wantErr {
            t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
        }
    })
}
```