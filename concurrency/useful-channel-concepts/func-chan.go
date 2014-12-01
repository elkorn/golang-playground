package main

import "fmt"

func abstractListener(funChan chan func() string) {
	funChan <- func() string {
		return "Sent data."
	}
}

func showFuncChannel() {
	funChan := make(chan func() string)
	go abstractListener(funChan)
	select {
	case response := <-funChan:
		fmt.Println("Received:")
		fmt.Println(response())
	}

	close(funChan)
}
