package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}
}
