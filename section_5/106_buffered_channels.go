package main

import (
	"fmt"
	"time"
)

/*
Buffered Channels
- Why use buffered channels?
	- Asynchronous communication
	- Load balancing
	- Flow control
- Creating buffered channels
	- make(chan Type, capacity)
	- Buffer capacity
- Key concepts of channel buffering
	- Blocking behavior
	- Non-blocking operations
	- Impact on performance
- Best practices for using buffered channels
	- Avoid over-bufferening
	- Graceful shutdown
	- Monitor buffer usage
*/

// func main() {
// 	// make(chan Type, capacity)
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3 // at this point there should be a receiver or the program will panic as the channel has reached capacity
// 	fmt.Println("Buffered channels")
// }

// func main() {
// 	// make(chan Type, capacity)
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	ch <- 3
// 	fmt.Println("Buffered channels")
// }

// Blocking on send only if the buffer is full
// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Received: ", <-ch)
// 	}()
// 	fmt.Println("Blocking starts")
// 	ch <- 3 // blocks because the buffer is full
// 	fmt.Println("Blocking ends")
// 	fmt.Println("Received: ", <-ch)
// 	fmt.Println("Received: ", <-ch)
// 	fmt.Println("Buffered channels")
// }

// Blocking on receive only only if the buffer is empty
func main() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("End of program.")
}
