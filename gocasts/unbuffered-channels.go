package main

import (
	"fmt"
	"time"
)

func longRoutine(done chan<- bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func main() {
	ok := make(chan bool)
	go longRoutine(ok)
	fmt.Println("Start waiting")
	<-ok
	fmt.Println("Done waiting")
}
