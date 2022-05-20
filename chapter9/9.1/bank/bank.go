package bank

import (
	"errors"
)

var deposits = make(chan int)       // send amount to deposit
var balances = make(chan int)       // receive balance
var withdraws = make(chan withdraw) // withdraw fund

type withdraw struct {
	amount int
	err    chan error
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	errCh := make(chan error)

	w := withdraw{
		amount: amount,
		err:    errCh,
	}

	withdraws <- w
	if err := <-errCh; err != nil {
		return false
	}

	return true
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdraws:
			if balance < w.amount {
				w.err <- errors.New("insufficient fund")
				break
			}

			balance -= w.amount
			close(w.err)
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
