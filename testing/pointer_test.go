package testing_test

import (
	"errors"
	"fmt"
	"testing"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := new(Wallet)
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := &Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30)

		assertBalance(t, &wallet, Bitcoin(20))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func assertBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertError(t *testing.T, err error, want string) {
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
	if err.Error() != want {
		t.Errorf("got '%s', want '%s'", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("wanted an error but didn't get one")
	}
}
