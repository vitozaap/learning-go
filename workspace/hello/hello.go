package main

import (
	"example/reversed"
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	message := reverse.String("Hello World!")
	reversed, _ := reversed.ReverseInt(10)
	fmt.Println(reversed)

	fmt.Print(message)
}
