package main

import (
	"fmt"
	"github.com/mergermarket/go-accounting/bank"
	"sync"
)

var workerCount = 1000

func main() {
	cjBank := bank.ChrisBank{}
	bankMe(&cjBank)
}

func bankMe(account bank.Account) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			account.Deposit(10)
			wg.Done()
		}()
	}

	wg.Wait()

	endBalance := account.Withdraw(1)
	fmt.Println("End balance should be 9999 -> ", endBalance)
}
