# Testing in Go

Go provides a built-in testing framework that makes it easy to write and run tests.

## Basic Testing

### Test File Naming
- Test files end in `_test.go`
- In the same package as the code being tested
- Excluded from regular builds

### Test Function Naming
```go
func TestMyFunction(t *testing.T)        // Unit test
func BenchmarkMyFunction(b *testing.B)   // Benchmark
func ExampleMyFunction()                 // Example
```

## Types of Tests

### 1. Unit Tests
```go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

### 2. Table-Driven Tests
```go
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 6},
        {"negative", -2, 3, -6},
        {"zero", 0, 5, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Multiply(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("got %d, want %d", got, tt.expected)
            }
        })
    }
}
```

### 3. Benchmarks
```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(10)
    }
}
```

### 4. Examples
```go
func ExampleHello() {
    fmt.Println(Hello("world"))
    // Output: Hello, world!
}
```

## Testing Tools

### 1. Test Helpers
```go
func setupTestCase(t *testing.T) func() {
    t.Helper() // Marks this as a helper function
    
    // Setup code
    return func() {
        // Teardown code
    }
}
```

### 2. Test Coverage
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 3. Race Detector
```bash
go test -race ./...
```

## Best Practices

1. **Test Organization**
   - Group related tests together
   - Use descriptive test names
   - Keep tests focused and simple
   - Follow AAA pattern (Arrange, Act, Assert)

2. **Test Coverage**
   - Aim for high test coverage
   - Test edge cases
   - Test error conditions
   - Don't test just for coverage

3. **Test Independence**
   - Tests should not depend on each other
   - Clean up after tests
   - Use setup and teardown helpers
   - Avoid global state

4. **Test Readability**
   - Write clear test descriptions
   - Use table-driven tests for multiple cases
   - Keep assertions simple
   - Comment complex test logic

## Common Patterns

### 1. Setup and Teardown
```go
func TestMain(m *testing.M) {
    // Setup
    code := m.Run()
    // Teardown
    os.Exit(code)
}
```

### 2. Mocking
```go
type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) Get(id string) (string, error) {
    args := m.Called(id)
    return args.String(0), args.Error(1)
}
```

### 3. Subtests
```go
t.Run("subtest name", func(t *testing.T) {
    // Subtest code
})
```

## Advanced Testing

### 1. Testing HTTP Handlers
```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    handler(w, req)
    
    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("got status %d, want %d", 
            resp.StatusCode, http.StatusOK)
    }
}
```

### 2. Testing with Context
```go
func TestWithContext(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    if err := functionWithContext(ctx); err != nil {
        t.Error(err)
    }
}
```

### 3. Parallel Testing
```go
func TestParallel(t *testing.T) {
    t.Parallel()
    // Test code
}
```

## Testing Commands

```bash
# Run all tests
go test ./...

# Run specific test
go test -run TestName

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=.

# Run tests with coverage
go test -cover ./...

# Generate HTML coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests with race detector
go test -race ./...
```