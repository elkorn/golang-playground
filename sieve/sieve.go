package main

import "log"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	log.Printf("Current prime: %d\n", prime)
	for {
		i := <-in
		// If a candidate is not divisivle by given prime, it can be passed along for further filtering.
		if i%prime != 0 {
			// log.Printf("New candidate %d...\n", i)
			// log.Printf("%d %% %d\t-> %v", i, prime, bool(i%prime != 0))
			out <- i
		}
	}
}

func main() {
	count := 50
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < count; i++ {
		prime := <-ch
		log.Println()
		log.Printf("%d\n", prime)
		log.Println()
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		// Go forward in the daisy-chain.
		// Make the old output the new input + make the old output the holder of the new prime.
		ch = ch1
	}
}
