package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var balance int
var transactionNo int
var transactions int

func synchronizedByChannels(start, done func()) {
	balanceChan := make(chan int)
	transactionChan := make(chan bool)

	start()
	for i := 0; i < transactions; i++ {
		go func(i int) {
			transactionAmount := rand.Intn(25)
			balanceChan <- transactionAmount
			if i == transactions-1 {
				// fmt.Println("Should quit")
				transactionChan <- true
				close(balanceChan)
				done()
			}
		}(i)
	}

	shouldContinue := true

	// go transaction(0)
	for shouldContinue {
		select {
		case amt := <-balanceChan:
			transaction(amt)
		case status := <-transactionChan:
			if status {
				// fmt.Println("Done")
				shouldContinue = false
				close(transactionChan)
			}
		}
	}
}

func withRaceCondition(start, done func()) {
	transactionChan := make(chan bool)

	start()
	for i := 0; i < transactions; i++ {
		go func(i int, transactionChan chan bool) {
			transactionAmount := rand.Intn(25)
			transaction(transactionAmount)
			if i == transactions-1 {
				// fmt.Println("Should quit")
				transactionChan <- true
			}
		}(i, transactionChan)
	}

	// go transaction(0)
	select {
	case <-transactionChan:
		// fmt.Println("Transactions finished")
		done()
	}

	close(transactionChan)
}

func initialize() {
	balance = 1000
	transactions = 100
	transactionNo = 0
}

func run(action func(func(), func())) {
	initialize()
	var wg sync.WaitGroup
	done := func() {
		wg.Done()
	}
	start := func() {
		wg.Add(1)
	}

	// fmt.Println("Starting balance: $", balance)
	action(start, done)
	wg.Wait()
	// fmt.Println("Final balance: $", balance)
}

func transaction(amount int) bool {
	approved := false
	// Maintain the non-negative acct balance invariant.
	if balance-amount >= 0 {
		approved = true
		balance -= amount
	}

	// var approvedText string
	// if approved {
	// 	approvedText = "approved"
	// } else {
	// 	approvedText = "declined"
	// }

	transactionNo++
	// fmt.Println(transactionNo, "transaction for $", amount, approvedText)
	// fmt.Println("\t Remaining balance $", balance)
	return approved
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	run(synchronizedByChannels)
	// run(withRaceCondition)
}
