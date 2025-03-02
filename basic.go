package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("it's super boring")
	fmt.Println("I'm leaving")
}

// this function will not be called because the main thread is main exit first
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
