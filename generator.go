package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator pattern is a pattern that will return a channel
// the channel will be used to communicate between goroutines
// the goroutine that will send data to the channel will be created inside the function
// the caller will receive the channel and use it to receive data from the channel
// the caller will be blocked until the sender sends data to the channel
// the sender will be blocked until the caller receives data from the channel
// this pattern is useful when we want to create a generator that will be used by multiple goroutines
func main() {
	go boringReturnChannel("it's super boring")
	c := boringReturnChannel("boring")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("I'm leaving")
}

func boringReturnChannel(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			fmt.Sprintf("%s %d", msg, i)
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
