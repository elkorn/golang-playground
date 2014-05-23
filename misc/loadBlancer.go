package main

import (
	"container/heap"
	"math/rand"
	"time"
)

type Request struct {
	operation     func() int
	resultChannel chan int
}

type Worker struct {
	requests chan Request // work to do (buff'd chan)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Pop() *Worker {
	res := p[p.Len()]
	p = Pool(p[0:p.Len()])
	return res
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	w.pending--
	// Change the place of the worker within the heap,
	// according to its new load.
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.completed(w)
		}
	}
}

func furtherProcess(result int) {
	/* ... */
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.resultChannel <- req.operation()
		done <- w
	}
}

// simulation of a requester
func requester(work chan Request) {
	c := make(chan int)

	for {
		time.Sleep(rand.Int63n(nWorker * 2e9))
		work <- Request{workFn, c}
		result := <-c
		furtherProcess(result)
	}
}

func main() {

}
