package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("boring")
	time.Sleep(time.Second * 1)
	fmt.Println("i'm leaving")
}

func boringWithSleep(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3) * int(time.Millisecond)))
	}
}
