package main

import "fmt"

// i'm tired...

func sum(a []int, c chan int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	res := 0
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, "+", y, "=", x+y)
	for {
		x, ok := <-c
		res += x
		if !ok {
			break
		}
	}

	fmt.Println(res)
}
