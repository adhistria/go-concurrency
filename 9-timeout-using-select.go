package main

import (
	"fmt"
	"math/rand"
	"time"
)

// time.After is returning channel that will send data after the given time
// this is useful to create timeout mechanism
// because we declare time.After in select method
// if the channel return first, we will get the data, and time.After will be reset
// if the time.After return first, we will get the timeout message and return
func main() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)

		case <-time.After(800 * time.Millisecond):
			fmt.Println("You are too slow")
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
