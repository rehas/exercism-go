package account

import "sync"

// Define the Account type here.

type Account struct {
	mu      *sync.Mutex
	closed  bool
	balance int64
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{closed: false, balance: amount, mu: &sync.Mutex{}}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	if amount < 0 { // withdrawal, check enough balance
		if a.balance < -1*amount {
			return 0, false // not enough balance
		}
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	remaining := a.balance
	a.balance = 0
	a.closed = true

	return remaining, true
}
