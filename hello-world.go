package main

import (
	"fmt"
	// "rsc.io/quote"
)

// Fuc returns a string
func Hello(name string) string {
	//a function whose name starts with a capital letter
	message := fmt.Sprint("Hi, %v. Welcome!", name)
	return message
}
func main() {
	// fmt.Println(quote.Go())

	var message string
	Hello(message)
}
