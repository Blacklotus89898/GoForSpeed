package main // Package declaration required for every Go executable

import (
	"fmt"
	"github.com/google/uuid"
	"go-for-speed/src/sandbox/anotherPackage" // Importing a custom package, relative to $GOPATH
	"time"                              
	"runtime"
)

// Interface
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func main() {
	fmt.Println("Hello, World!")

	var name string = "Go"
	age := 10 // Short declaration
	fmt.Println("Name is", name, "and age is", age)

	id := uuid.New()
	fmt.Println("Generated UUID:", id)

	var result int = util.Double(10)
	fmt.Println("Result:", result)

	// Data types
	var a int = 10       // Integer
	var b float64 = 3.14 // Float
	var c bool = true    // Boolean
	var d string = "Go"  // String
	fmt.Println(a, b, c, d)

	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Running")
		break
	}

	// Loop over slices/maps
	arr := []string{"A", "B", "C"}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println(add(10, 20))

	quotient, ok := divide(10, 0)
	if ok {
		fmt.Println("Quotient:", quotient)
	} else {
		fmt.Println("Division by zero")
	}

	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Name) // Outputs: Alice

	// Concurrency
	go func() {
		fmt.Println("Running in a goroutine")
	}()
	fmt.Println("Running in main routine")

	// Channels
	messages := make(chan string)
	go func() { messages <- "Hello from goroutine" }()
	fmt.Println(<-messages)

	err := readFile("example.txt")
	if err != nil {
		fmt.Println(err)
	}

	var x int = 10
	var ptr *int = &x // Pointer to x
	fmt.Println(*ptr) // Dereference

	// Methods on structs
	p.greet()

	// Goroutine
	go sayHello()           // Runs concurrently
	time.Sleep(time.Second) // Prevents main from exiting immediately

	// Interface
	var animal Animal = Dog{}
	fmt.Println(animal.Speak())

	// Channels
	ch := make(chan int) // Create a channel

	go func() {
		ch <- 42 // Send data into the channel
	}()

	val := <-ch // Receive data from the channel
	fmt.Println("Received:", val)

	buffered_ch := make(chan int, 2) // Buffered channel with capacity 2
	buffered_ch <- 1                 // Non-blocking, space available in buffer
	buffered_ch <- 2                 // Non-blocking, fills buffer
	close(ch)

	val, ok = <-ch
	if !ok {
		fmt.Println("Channel closed!")
	} else {
		fmt.Println("Received:", val)
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println("Memory Used:", memStats.Alloc)

}

// Method for the Person struct, like implementing a trait in rust
func (p Person) greet() {
	fmt.Println("Hello,", p.Name)

}

func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Struct
type Person struct {
	Name string
	Age  int
}

// Error handling
func readFile(filename string) error {
	return fmt.Errorf("failed to read file: %s", filename)
}

func sayHello() {
	fmt.Println("Hello")
}

func Add(a, b int) int {
	return a + b
}
