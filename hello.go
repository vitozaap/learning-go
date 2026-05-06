package main

import "fmt"
import "rsc.io/quote"
import "github.com/docker/go-units"

func main() {
	var name string = "Jonh Doe"
	fmt.Printf("Hello %s!", name)
	newQuote()
	fmt.Println(bytesConverter(2048))
}

func bytesConverter(bytes int32) string {
	return units.BytesSize(float64(bytes))
}

func newQuote() {
	fmt.Println(quote.Go())
}
