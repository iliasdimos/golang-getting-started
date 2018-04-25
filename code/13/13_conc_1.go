package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("goroutine") // concurrently

	f("direct") // synchronously

	// You can also start a goroutine for an anonymous func call
	go func(msg string) { fmt.Println(msg) }("on going")

	// We use fmt.Scanln or sleep to block here
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
