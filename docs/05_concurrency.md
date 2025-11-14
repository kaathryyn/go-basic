# Concurrency in Go

Go provides powerful built-in features for concurrent programming through goroutines and channels.

## Goroutines

Goroutines are lightweight threads managed by the Go runtime.

### Basic Goroutine Usage
```go
go function() // Start a new goroutine
```

### Example:
```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func main() {
    go printNumbers() // Run in background
    time.Sleep(time.Second) // Wait for goroutine
}
```

## Channels

Channels are the pipes that connect concurrent goroutines.

### Channel Types
```go
ch := make(chan int)    // Unbuffered channel
ch := make(chan int, 5) // Buffered channel with capacity 5
```

### Channel Operations
```go
ch <- x    // Send x to channel ch
x := <-ch  // Receive from channel ch
close(ch)  // Close channel
```

### Directional Channels
```go
chan<- int // Send-only channel
<-chan int // Receive-only channel
```

## Synchronization

### WaitGroup
WaitGroup is used to wait for a collection of goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)    // Add a counter
wg.Done()    // Decrease counter
wg.Wait()    // Block until counter is zero
```

### Mutex
Mutex is used to protect shared resources.

```go
var mu sync.Mutex

mu.Lock()    // Lock the mutex
// Critical section
mu.Unlock()  // Unlock the mutex
```

## Select Statement

Select lets you wait on multiple channel operations.

```go
select {
case msg1 := <-ch1:
    // Use msg1
case ch2 <- x:
    // Sent x to ch2
case <-time.After(time.Second):
    // Timeout after 1 second
default:
    // Run if no other case is ready
}
```

## Best Practices

### 1. Goroutine Management
- Don't create goroutines in libraries
- Make it clear when goroutines start and stop
- Check for goroutine leaks
- Use context for cancellation

### 2. Channel Usage
- Use channels to communicate between goroutines
- Use buffered channels when appropriate
- Always handle channel closure
- Document channel ownership

### 3. Error Handling
- Propagate errors through channels
- Use error channels for concurrent error handling
- Consider using errgroup for concurrent tasks

### 4. Resource Protection
- Use mutex for simple shared resource protection
- Consider using atomic operations for counters
- Use sync.Once for one-time initialization

## Common Patterns

### 1. Worker Pool
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}
```

### 2. Pipeline
```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
```

### 3. Fan-out, Fan-in
```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            for n := range ch {
                out <- n
            }
            wg.Done()
        }(c)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

## Testing Concurrent Code

### 1. Race Detector
```bash
go test -race ./...
```

### 2. Timing-Dependent Tests
```go
func TestConcurrent(t *testing.T) {
    done := make(chan bool)
    go func() {
        // Do work
        done <- true
    }()
    
    select {
    case <-done:
        // Test passed
    case <-time.After(time.Second):
        t.Error("timeout")
    }
}
```

### 3. Channel Testing
```go
func TestChannel(t *testing.T) {
    ch := make(chan int)
    go func() {
        ch <- 42
        close(ch)
    }()
    
    got := <-ch
    want := 42
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```