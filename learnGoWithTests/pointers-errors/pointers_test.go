package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		checkBalance(expected, t, wallet)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(5))
		expected := Bitcoin(5)

		checkNoError(t, err)
		checkBalance(expected, t, wallet)
	})

	t.Run("Withdraw Insufficient Funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkError(t, err, ErrInsuffientFunds)
		checkBalance(startingBalance, t, wallet)
	})
}

func checkBalance(expected Bitcoin, t testing.TB, wallet Wallet) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("Result %s | Expected %s", result.String(), expected.String())
	}
}

func checkError(t testing.TB, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("Expected error but didn't get one")
	}

	if result != expected {
		t.Errorf("Result %q | Expected %q", result, expected)
	}
}

func checkNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("expected no error, but got one: ", err)
	}
}
