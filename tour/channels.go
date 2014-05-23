package main

import (
	"fmt"
	"time"
)

// Fixed-sized buffer channels

func fixedSizeBuffer() {
	// add a superfluous read/write and see what happens.
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i += 1 {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func runEx(ex func()) {
	ex()
	fmt.Println()
}

func selectStatement() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacciWithSelect(c, quit)
}

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func defaultCase() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print("...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	runEx(fixedSizeBuffer)
	runEx(rangeAndClose)
	runEx(selectStatement)
	runEx(defaultCase)
}
