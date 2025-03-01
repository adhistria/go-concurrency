package main

import "fmt"


func main() {
	go boring("it's super boring")
	fmt.Println("I'm leaving")
}

// this function will not be called because the main thread is main exit first
func boring(msg string) {
	fmt.Println(msg)	
}