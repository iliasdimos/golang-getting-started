package main

import "fmt"

func main() {
	defer fmt.Println("End")
	go fmt.Println("hello world")
}
