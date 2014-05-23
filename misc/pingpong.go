package main

import (
	"fmt"
	"time"
)

// godoc.org -> a generated documentation for Go packages.
type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("pong", table)
	go player("ping", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table // grab the ball after the timeout
	// The stack traces show that the goroutines have leaked.
	// Blocking profiles - display graphs.
	panic("gimme stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table // Wait for the ball to arrive.
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
