package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I'm listening.")
	count := 5
	fmt.Println(count)
	//go boring("msg")

	// ch1 := generatorOfBoring("generated boring 1")
	// ch2 := generatorOfBoring("generated boring 2")
	// for i := 0; i < count; i++ {
	//  fmt.Println(<-ch1)
	//  fmt.Println(<-ch2)
	// }

	// ch := fanIn(generatorOfBoring("fanned1"), generatorOfBoring("fanned2"))
	// for i := 0; i < count; i++ {
	//  fmt.Println(<-ch)
	// }

	// sequenced execution of concurrent procedures
	// ch := seqFanIn(generatorOfBoring("fanned1"), generatorOfBoring("fanned2"))
	// for i := 0; i < count; i++ {
	//  msg1 := <-ch
	//  fmt.Println(msg1.str)
	//  msg2 := <-ch
	//  fmt.Println(msg2.str)
	//  msg1.wait <- true
	//  msg2.wait <- true
	// }

	// select used for sequencing communication
	// ch := selectedFanIn(generatorOfBoring("fanned1"), generatorOfBoring("fanned2"))
	// for i := 0; i < count; i++ {
	// 	fmt.Println(<-ch)
	// }

	// // timing out each message
	// timeout()

	// // timing out the whole conversation
	// totalTimeout()

	// // using an externally controlled quit channel to end goroutines
	// quitChannel()

	// chinese whispers game simulation
	daisyChain()
	fmt.Println("You're boring!")
}

// Channels communicate AND synchronize when sending/receiving values.

func boring(msg string) {
	for i := 0; ; i++ {
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		fmt.Println(msg, i)
	}
}

func generatorOfBoring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1
		}
	}()
	go func() {
		for {
			out <- <-input2
		}
	}()
	// fmt.Println(len(inputs))
	// for _, c := range inputs {
	// 	go func() {
	// 		for {
	// 			out <- <-c
	// 		}
	// 	}()
	// }

	return out
}

func seqFanIn(input1, input2 <-chan string) <-chan Message {
	out := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for {
			out <- Message{<-input1, waitForIt}
			<-waitForIt
		}
	}()
	go func() {
		for {
			out <- Message{<-input2, waitForIt}
			<-waitForIt
		}
	}()

	return out
}

type Message struct {
	str  string
	wait chan bool
}

func selectedFanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
				// default:
				// 	c <- fmt.Sprint("...")
			}
		}
	}()

	return c
}

func timeout() {
	c := generatorOfBoring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func totalTimeout() {
	c := generatorOfBoring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boringWithQuit(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// noop
			case <-quit:
				return
			}
		}
	}()

	return c
}

func quitChannel() {
	quit := make(chan bool)
	c := boringWithQuit("Joe", quit)
	x := rand.Intn(10)
	for i := x; i > 0; i-- {
		fmt.Println(<-c)
	}

	fmt.Println("After", x)

	quit <- true
}

func daisyChain() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go daisyLink(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func daisyLink(left, right chan int) {
	// When a value is being passed from the right channel, it's being incremented and passed on to the left.
	left <- 1 + <-right
}
