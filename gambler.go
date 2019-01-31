package gambler

import (
	"fmt"
)

// WinThreshold sets the number you have to match or beat to win
const WinThreshold = 80

// A GamblingSession is a record about a gambling session
type GamblingSession struct { 
	attempts int
	wins int
	losses int
	balance int
}

// Attempts returns the number of attempts left in the current session
func (gs *GamblingSession) Attempts() int {
	return gs.attempts
}

// Wins returns the number of wins in the current session
func (gs *GamblingSession) Wins() int {
	return gs.wins
}

// Losses returns the number of losses in the current session
func (gs *GamblingSession) Losses() int {
	return gs.losses
}

// Balance returns the remaining balance in the current session
func (gs *GamblingSession) Balance() int {
	return gs.balance
}

// BuyAttempts reduces the balance and adds attempts to the current session
func (gs *GamblingSession) BuyAttempts(quantity int) error {
	//todo implement BuyAttempts and validate input
	if quantity <= 0 {
		return fmt.Errorf("%d is not a positive integer", quantity)
	} else if quantity > gs.balance {
		return fmt.Errorf("%d attempts costs more than the balance of %d", quantity, gs.balance)
	} else {
		gs.balance -= quantity
		gs.attempts += quantity
	}
	return nil
}

// MakeAnAttempt generates a random number and checks it against the WinThreshold
func (gs *GamblingSession) MakeAnAttempt(roll int) (bool, int, error) {
	//todo implement MakeAnAttempt
	var result bool 
	
	if roll >= WinThreshold {
		result = true
		gs.wins++
	} else {
		result = false
		gs.losses++
	}
	gs.attempts--
	return result, roll, nil
}

// AddToBalance adds money to the balance of the current session
func (gs *GamblingSession) AddToBalance(dollars int) error {
	//todo validate input
	if dollars <= 0 {
		return fmt.Errorf("%d is not a positive integer", dollars)
	}
	gs.balance += dollars
	return nil
}

// CashOut refunds attempts left in the current session
func (gs *GamblingSession) CashOut() error {
	//todo implement CashOut
	gs.balance += gs.attempts
	gs.attempts = 0
	return nil
}

