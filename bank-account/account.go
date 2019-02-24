package account

type AccountFunctions interface {
	Open(initialDeposit int64) *Account
	Close() (payout int64, ok bool)
}

type Account struct {
	AccountValid bool
	Amount       int64
}

func Open(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		AccountValid: true,
		Amount:       initialDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {

	if a.AccountValid == true {
		a.AccountValid = false
		return a.Amount, !a.AccountValid
	}

	return a.Amount, a.AccountValid
}

func (a *Account) Balance() (balance int64, ok bool) {

	if a.AccountValid == false {
		return 0, a.AccountValid
	}

	return a.Amount, a.AccountValid
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	if a.AccountValid == false {
		return a.Amount, false
	} else if a.Amount+amount < 0 {
		return amount, false
	}

	a.Amount = a.Amount + amount

	return a.Amount, true
}
