package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Example 1: Goroutines
	go sayHello()
	time.Sleep(1 * time.Second) // Wait for goroutine to finish

	// Example 2: Channels
	ch := make(chan string)
	go sendMessage(ch) // Send message to channel
	message := <-ch
	fmt.Println("Received from channel:", message)

	// Example 3: Buffered Channels
	bufferedCh := make(chan int, 2)
	bufferedCh <- 1
	bufferedCh <- 2
	fmt.Println("Buffered channel values:", <-bufferedCh, <-bufferedCh)

	// Example 4: WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 finished")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 finished")
	}()
	wg.Wait()

	// Example 5: Select Statement
	selectCh1 := make(chan string)
	selectCh2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		selectCh1 <- "Message from channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		selectCh2 <- "Message from channel 2"
	}()
	select {
	case msg1 := <-selectCh1:
		fmt.Println(msg1)
	case msg2 := <-selectCh2:
		fmt.Println(msg2)
	}

	fmt.Println("All examples completed!")
}

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func sendMessage(ch chan string) {
	ch <- "Hello, Channel!"
}