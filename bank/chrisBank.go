package bank

import "sync"

type ChrisBank struct {
	balance int
	mutex   sync.Mutex
}

func (c *ChrisBank) addToBalance(amount int) (newBalance int) {
	c.mutex.Lock()
	c.balance += amount
	c.mutex.Unlock()
	return c.balance
}

func (c *ChrisBank) Withdraw(amount int) int {
	return c.addToBalance(-amount)
}

func (c *ChrisBank) Deposit(amount int) int {
	return c.addToBalance(amount)
}