package account

import "sync"

type Account struct {
	money  int64
	opened bool
	mu     sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{money: amount, opened: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.opened {
		return a.money, true
	}

	return 0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened {
		return 0, false
	}
	if amount < 0 && a.money < 0-amount {
		return 0, false
	}
	a.money += amount

	return a.money, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened {
		return 0, false
	}
	a.opened = false
	return a.money, true
}
