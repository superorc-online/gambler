package main

import "fmt"
import "math/rand"
import "time"
import "bufio"
import "os"
import "strings"

func main() {

var attempts int
var playtime int = 1
var outofTurns int = 1
var winner int = 1
var attempting int = 1


rando := rand.NewSource(time.Now().UnixNano())
reader := bufio.NewReader(os.Stdin)
scanner := bufio.NewScanner(os.Stdin)


	fmt.Printf("Welcome to The Random Number Draw! To win you have to draw an 80 or higher. Each try cost $1.00. Would you like to play? (Yes/No) ")
Full:
for (playtime != 0){

		scanner.Scan()
		play := scanner.Text()
		play = strings.TrimRight(play, "\r\n")
		play = strings.ToUpper(play)
		
		if (play == "YES"){
			Play:
			for (playtime !=0){
				for (attempts <= 0) {
					fmt.Printf("How many attempts would you like to purchase? ")
					fmt.Scan(&attempts)
				}
				
			
				fmt.Printf("\nGreat! You have purchased %d attempts for $%d.00.\nTime to play!\n\n", attempts, attempts)
			Cont:
			for (playtime != 0){
				random := rand.New(rando)
				var n = random.Intn(100)

				if(n >= 80){
					fmt.Printf("Your number is %d which means you WIN!\n", n)
					attempts --
						if (attempts == 0){
							fmt.Printf("\nWould you like to play again? (Yes/No) ")
							for (winner != 0){
							scanner.Scan()
							playAgain := scanner.Text()
							playAgain = strings.TrimRight(playAgain, "\r\n")
							playAgain = strings.ToUpper(playAgain)
								if (playAgain == "YES"){
									continue Play
								} else if (playAgain == "NO") {
									playtime = 0
									break
								} else if (playAgain == ""){
									fmt.Printf("You input nothing. Please enter Yes or No. ")
								} else {
									fmt.Printf("Your input is not valid. Please enter Yes or No. ")
								}
							}
						} else {
							fmt.Printf("You have %d attempts remaining. Would you like to continue playing? (Yes/No) ", attempts)
							for (attempting != 0){
							scanner.Scan()
							cont := scanner.Text()
							cont = strings.TrimRight(cont, "\r\n")
							cont = strings.ToUpper(cont)
							if (cont == "YES") {
								continue Cont
							} else if (cont == "NO") {
								fmt.Printf("You have been refuned $%d.00. Thanks for playing!\nGoodbye.", attempts)
								playtime = 0
								continue Full
							} else if (cont == ""){
								fmt.Printf("You input nothing. Please enter Yes or No. ")
							} else {
								fmt.Printf("Your input is not valid. Please enter Yes or No. ")
							}
							}
						}
				} else{
					fmt.Printf("Your number is %d which means you lose.\n", n)
					attempts --
					if (attempts != 0) {
						break
					} else if (attempts == 0) {
						fmt.Printf("Would you like to play again? (Yes/No) ")
					for (outofTurns != 0){
						scanner.Scan()
						outOfTurns := scanner.Text()
						outOfTurns = strings.TrimRight(outOfTurns, "\r\n")
						outOfTurns = strings.ToUpper(outOfTurns)
						if (outOfTurns == "YES") {
							continue Play
						} else if (outOfTurns == "NO") {
							playtime = 0
							fmt.Printf("Goodbye.")
							continue Full
						} else if (outOfTurns == ""){
							fmt.Printf("You input nothing. Please enter Yes or No. ")
						} else {
							fmt.Printf("Your input is not valid. Please enter Yes or No. ")
						}
					}
					}
				}
					attempts --
						if (attempts == 0){
							fmt.Printf("\nWould you like to play again? (Yes/No) ")
							for (attempting != 0) {
							scanner.Scan()
							playAgain := scanner.Text()
							playAgain = strings.TrimRight(playAgain, "\r\n")
							playAgain = strings.ToUpper(playAgain)
								if (playAgain == "YES"){
									continue Play
								} else if (playAgain == "NO") {
									playtime = 0
								} else if (playAgain == ""){
									fmt.Printf("You input nothing. Please enter Yes or No. ")
								} else {
									fmt.Printf("Your input is not valid. Please enter Yes or No. ")
								}
							}
						}
			}
				Loser:
				for (attempts != 0){
					fmt.Printf("You have %d tries left to get an 80 or higher. Would you like to try again? (Yes/No) ", attempts)
					for (attempting != 0){
					input, _ := reader.ReadString('\n')
					input = strings.TrimRight(input, "\r\n")
					input = strings.ToUpper(input)
				if (input == "YES") {
					attempts --
					random := rand.New(rando)
					var n = random.Intn(100)
					if(n >= 80){
						fmt.Printf("Your number %d which means you WIN!\nGoodbye.", n)
						playtime = 0
						continue Full
					} else{
						fmt.Printf("Your number is %d which means you lose.\n", n)
						continue Loser
					}
				} else if (input == "NO") {
					fmt.Printf("You have been refunded $%d.00. Thanks for playing!\nGoodbye.", attempts)
					playtime = 0
					continue Full
				} else if (input == ""){
					fmt.Printf("You input nothing. Please enter Yes or No. ")
				} else {
					fmt.Printf("Your input, %s, is not valid. Please enter Yes or No. ", input)
				}
					}
					playtime = 0
				}
		}

		fmt.Println("\nGoodbye.")	

		} else if (play == "NO"){
			fmt.Println("Goodbye.")
			playtime = 0
		} else if (play == ""){
			fmt.Printf("You said nothing. Would you like to play? (Yes/No) ")
		} else {
			fmt.Printf("Your input, %s, is not valid. Please enter Yes or No. ", play)
		}
			

	

}



}	