package main

import (
	"fmt"
	"time"
)

// This example will show how to use generator pattern to create a lockstep goroutine
// even though we have 2 goroutines, the goroutines will wait for each other
// the goroutines will be synchronized
func main() {
	c1 := generator("Ann")
	c2 := generator("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(1e3) * time.Millisecond)
		}
	}()
	return c
}
