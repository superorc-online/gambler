package main

import (
	"fmt"
	
	"github.com/superorc-online/gambler"
)

func main() {
	gs := &gambler.GamblingSession{}
	gs.AddToBalance(5)
	fmt.Printf("Welcome to The Random Number Draw! To win you have to draw an %d or higher. Each try cost $1.00. Your current balance is $%d.00. Would you like to play? (Yes/No) ", gambler.WinThreshold, gs.Balance())
	//todo implement mainloop

	fmt.Println("Goodbye.")
}