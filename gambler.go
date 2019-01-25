package gambler

const WinThreshold = 80

type GamblingSession struct {
	attempts int
	wins int
	losses int
	balance int
}

func (gs *GamblingSession) Attempts() int {
	return gs.attempts
}

func (gs *GamblingSession) Wins() int {
	return gs.wins
}

func (gs *GamblingSession) Losses() int {
	return gs.losses
}

func (gs *GamblingSession) Balance() int {
	return gs.balance
}

func (gs *GamblingSession) BuyAttempts(quantity int) error {
	//todo implement BuyAttempts
	return nil
}

func (gs *GamblingSession) MakeAnAttempt() (bool, int, error) {
	//todo implement MakeAnAttempt
	return false, 0, nil
}

func (gs *GamblingSession) AddToBalance(dollars int) error {
	//todo validate input
	gs.balance += dollars
	return nil
}

func (gs *GamblingSession) CashOut() error {
	//todo implement CashOut
	return nil
}

