package pointerserrors

import "fmt"

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

type Stringer interface {
	String()
}

var ErrInsuffientFunds = fmt.Errorf("insufficient funds")

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsuffientFunds
	}
	wallet.balance -= amount
	return nil
}
