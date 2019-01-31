package gambler

import (
	"math/rand"
	"time"
	"fmt"
)

// WinThreshold sets the number you have to match or beat to win
const WinThreshold = 80

// GamblingSession todo comment
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
	if quantity > 0 {
		gs.balance -= quantity
		gs.attempts += quantity
	} else {
		return fmt.Errorf("%d is not a positive integer, please enter a positive integer", quantity)
	}
	return nil
}

// MakeAnAttempt generates a random number and checks it against the WinThreshold
func (gs *GamblingSession) MakeAnAttempt() (bool, int, error) {
	//todo implement MakeAnAttempt
	var result bool 
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(100)
	if roll >= WinThreshold {
		result = true
		gs.wins ++
	} else {
		result = false
		gs.losses ++
	}
	return result, roll, nil
}

// AddToBalance adds money to the balance of the current session
func (gs *GamblingSession) AddToBalance(dollars int) error {
	//todo validate input
	if dollars > 0 {
		gs.balance += dollars
	} else {
		return fmt.Errorf("%d is not a positive integer, please enter a positive integer", dollars)
	}
	return nil
}

// CashOut refunds attempts left in the current session
func (gs *GamblingSession) CashOut() error {
	//todo implement CashOut
	gs.balance += gs.attempts
	return nil
}

