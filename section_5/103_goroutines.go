package main

import (
	"fmt"
	"time"
)

/*
Notes:
- Goroutine scheduling in Go
	- Managed by the Go Runtime Scheduler
	- Uses M:N scheduling Model
	- Efficient multiplexing
- Common pitfalls and best practices
	- Avoiding goroutine leaks
	- Limiting goroutine creation
	- Proper error handling
	- Synchronization
*/

// Gorountines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished or ready to return any value.
// Goroutines do not stop the program flow and are non-blocking

func main() {
	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()
	// err = go doWork() // This is not accepted
	go printNumbers()
	go printLetters()

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// } else {
	// 	fmt.Println("Work completed successfully")
	// }

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

func sayHello() {
	time.Sleep(time.Second)
	fmt.Println("Hello from gorountine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter: ", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(time.Second)

	return fmt.Errorf("An error occured in doWork")
}
