package account

import "sync"

type AccountFunctions interface {
	Open(initialDeposit int64) *Account
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
	Deposit(amount int64) (newBalance int64, ok bool)
}

type Account struct {
	AccountValid bool
	Amount       int64
	mux          sync.Mutex
}

// Open will open up a bank account with initial deposit
func Open(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		AccountValid: true,
		Amount:       initialDeposit}
}

// Close will invalidate the bank account and return the payout 0
func (a *Account) Close() (payout int64, ok bool) {

	//only 1 go routine can access this function one at at time
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.AccountValid == true {
		a.AccountValid = false
		return a.Amount, !a.AccountValid
	}

	return 0, a.AccountValid
}

// Balance will get the amount that the user has
func (a *Account) Balance() (balance int64, ok bool) {
	// since it is read only, it is not necessary to have a mutex on it
	if a.AccountValid == false {
		return 0, a.AccountValid
	}

	return a.Amount, a.AccountValid
}

// Deposit will put the amount inside the account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	//only 1 go routine can access this function one at at time
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.AccountValid == false {
		return a.Amount, false
	} else if a.Amount+amount < 0 {
		return amount, false
	}

	a.Amount = a.Amount + amount

	return a.Amount, true
}
