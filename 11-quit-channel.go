package main

import (
	"fmt"
)

func main() {
	fmt.Println("I'm listening")
	quit := make(chan bool)
	c := boring("Joe", quit)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			fmt.Println("I'ts time to quit")
			return
		}
	}
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
		quit <- true
	}()
	return c
}
