package main

import (
	"fmt"
	"time"
)

/*
Notes:
- Why use channels?
	- Enable safe and efficient communication between concurrent goroutines
	- Helo synchronize and manage the flow of data in concurrent programs
- Basics of channels
	- Creating channels (make(chan Type))
	- Sending and receiving data (<-)
	- Channel directions
		- send-only: `ch <- value`
		- receive-only: `value := <-ch`
- Common pitfalls and best practices
	- Avoid deadlocks
	- Avoiding unnecessary buffering
	- Channel directino
	- Graceful shutdown
	- Use `defer` for unlocking
*/

func main() {
	//variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking because it is continously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"

		for _, e := range "abcde" {
			greeting <- "Letter: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting // blocks forever but someone must send a value to it or there will be panic
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting // blocks forever but someone must send a value to it or there will be panic
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting // blocks forever but someone must send a value to it or there will be panic
	fmt.Println(receiver)
	receiver = <-greeting // blocks forever but someone must send a value to it or there will be panic
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(time.Second)
	fmt.Println("End of program.")
}
