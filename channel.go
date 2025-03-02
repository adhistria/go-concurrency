package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel is a way to communicate between goroutines, 1 channel will have 2 goroutines
// 1 goroutine will send data to the channel, 1 goroutine will receive data from the channel
// channel is a blocking operation, if the channel is empty, the receiver will be blocked until the sender sends data to the channel
// if the channel is full, the sender will be blocked until the receiver receives data from the channel

func main() {
	c := make(chan string)
	go boringWithChannel("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("I'm leaving")
}

func boringWithChannel(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}
