package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world")
	time.Sleep(5 * time.Second)
	done <- true
}
func main() {
	defer fmt.Println("End")

	// Create our done channel and start our goroutine
	done := make(chan bool)
	go hello(done)

	// Blocking, this will wait until something is received on the done channel
	<-done
}
