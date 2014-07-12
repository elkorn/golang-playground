package c164

type signal chan bool

type dch struct {
	req  signal
	data chan int
}

func makedch() *dch {
	result := new(dch)
	result.req = make(signal)
	result.data = make(chan int)
	return result
}

func put(U *dch, val int) {
	U.data <- val
}

func get(U *dch) int {
	return <-U.data
}

func generate(ch chan<- int) {
	for i := 3; ; i++ {
		ch <- i
	}
}
