package main

import (
	"fmt"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := 10; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("You are boring, I'm leaving")
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				// sometimes the println is not executed
				// because the main function is already finished
				// so the program is terminated
				fmt.Println("I'm tired of talking")
				return
			}
		}
	}()
	return c
}
