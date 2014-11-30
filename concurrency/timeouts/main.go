package main

import (
	"fmt"
	"time"
)

func main() {
	theChan := make(chan string, 1)
	var limit time.Duration = 3 * time.Second
	go func() {}()

	select {
	case <-time.After(limit):
		fmt.Println("Killing the goroutine through a buffered kill-chan after", limit)
		close(theChan)
	}
}
