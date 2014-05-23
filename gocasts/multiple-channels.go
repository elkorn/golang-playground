package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	quit := make(chan bool)

	go watcher(quit)

	var input string
	fmt.Scanf("%s", &input)
	log.Println("main quit")
	close(quit)
}

func watcher(quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		case <-time.After(1 * time.Second):
			log.Println("timer fired.")
		}
	}
}
