package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple goroutine example
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

// Channel examples
func generateNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send number to channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

func squares(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// WaitGroup example
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Mutex example
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("1. Basic Goroutines:")
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second) // Wait for goroutines to finish

	fmt.Println("\n\n2. Channels:")
	numbers := make(chan int)
	squares := make(chan int)

	// Start goroutines
	go generateNumbers(numbers)
	go squares(numbers, squares)

	// Read from squares channel
	fmt.Println("Squares of numbers:")
	for square := range squares {
		fmt.Printf("%d ", square)
	}

	fmt.Println("\n\n3. WaitGroup:")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go worker(i, &wg)
	}
	wg.Wait() // Wait for all workers to finish

	fmt.Println("\n4. Mutex:")
	counter := Counter{}
	var wg2 sync.WaitGroup

	// Launch 5 goroutines that increment the counter
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg2.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())

	// Select example
	fmt.Println("\n5. Select Statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}