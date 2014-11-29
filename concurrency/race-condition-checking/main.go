package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var balance int
var transactionNo int
var transactions int

func synchronizedByChannels(wg *sync.WaitGroup) {
	balanceChan := make(chan int)
	transactionChan := make(chan bool)

	for i := 0; i < transactions; i++ {
		go func(i int) {
			transactionAmount := rand.Intn(25)
			balanceChan <- transactionAmount
			if i == transactions-1 {
				fmt.Println("Should quit")
				transactionChan <- true
				close(balanceChan)
				(*wg).Done()
			}
		}(i)
	}

	shouldContinue := true

	go transaction(0)
	for shouldContinue {
		select {
		case amt := <-balanceChan:
			fmt.Println("Transaction for $", amt)
			// Maintain the non-negative acct balance invariant.
			if balance-amt < 0 {
				fmt.Println("Transaction failed.")
			} else {
				balance -= amt
				fmt.Println("Transaction succeeded.")
			}

			fmt.Println("Current balance: $", balance)

		case status := <-transactionChan:
			if status {
				fmt.Println("Done")
				shouldContinue = false
				close(transactionChan)
			}
		}
	}
}

func initialize() {
	balance = 1000
	transactions = 100
	transactionNo = 0
}

func run(action func(*sync.WaitGroup)) {
	initialize()
	var wg sync.WaitGroup
	fmt.Println("Starting balance: $", balance)
	wg.Add(1)
	action(&wg)
	wg.Wait()
	fmt.Println("Final balance: $", balance)
}

func transaction(amount int) bool {
	approved := false
	if balance-amount >= 0 {
		approved = true
		balance -= amount
	}

	var approvedText string
	if approved {
		approvedText = "approved"
	} else {
		approvedText = "declined"
	}

	transactionNo++
	fmt.Println(transactionNo, "transaction for $", amount, approvedText)
	fmt.Println("\t Remaining balance $", balance)
	return approved
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	run(synchronizedByChannels)
}
