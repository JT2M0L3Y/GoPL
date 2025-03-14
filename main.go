package main

import (
	"fmt"
	"os"
)

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
	fmt.Println("Done!")
	os.Exit(0)
}
