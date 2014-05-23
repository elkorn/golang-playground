package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()

	// Setting a channel to `nil` effectively turns it off in a `select` statement, so that it does not cause
	// unnecessary blockage when the program is supposed to wait for messages from it.
	if rand.Intn(2) == 0 {
		a = nil
		fmt.Println("nil a")
	} else {
		b = nil
		fmt.Println("nil b")
	}

	select {
	case s := <-a:
		fmt.Println("got", s)
	case s := <-b:
		fmt.Println("got", s)
	}

}
