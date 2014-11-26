package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string

var length int

func addToFinalStack(letters chan string, wg *sync.WaitGroup) {
	finalString += <-letters
	wg.Done()
}

func capitalize(letters chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letters <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	initialString = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

	initialBytes := []byte(initialString)

	var letters = make(chan string)

	length = len(initialString)

	for i := 0; i < length; i++ {
		wg.Add(2)
		go capitalize(letters, string(initialBytes[i]), &wg)
		go addToFinalStack(letters, &wg)

		wg.Wait()
	}

	fmt.Println(finalString)
}
