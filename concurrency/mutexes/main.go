package main

import (
	"fmt"
	"runtime"
	"sync"
)

func basicSync() {
	current := 0
	iterations := 100

	var wg sync.WaitGroup

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			current++
			fmt.Printf("%d ", current)
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("")
}

func withMutex() {
	current := 0
	iterations := 100

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			mutex.Lock()
			current++
			mutex.Unlock()
			fmt.Printf("%d ", current)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("")
}

func withoutMutex() {
	current := 0
	iterations := 100

	var wg sync.WaitGroup
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			current++
			fmt.Printf("%d ", current)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("")
}

func main() {
	runtime.GOMAXPROCS(2)
	basicSync()
	withoutMutex()
	withMutex()
}
