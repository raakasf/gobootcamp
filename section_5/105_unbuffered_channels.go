package main

import (
	"fmt"
	"time"
)

// This will panic as we need a channel as soon as we send to the channel
// func main() {
// 	ch := make(chan int)
// 	ch <- 1

// 	receiver := <-ch
// 	fmt.Println(receiver)
// }

// This will work as expected
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 1
// 	}()

// 	receiver := <-ch
// 	fmt.Println(receiver)
// }

// This will wait for timeout and the panic as the receiver will wait for all goroutines to exit (as they may send it a value) and then panic as it won't receive a value
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		// ch <- 1
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("2 second Goroutine finished")
// 	}()
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println("3 second Goroutine finished")
// 	}()
// 	receiver := <-ch
// 	fmt.Println(receiver)
// 	fmt.Println("End of program")
// }

// Unbuffered channels block on send as they wait for the goroutine to finish
func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)

		fmt.Println(<-ch)
		fmt.Println("3 second Goroutine finished")
	}()
	ch <- 1
	fmt.Println("End of program")
}
