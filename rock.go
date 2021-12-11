package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getGameNum() int {
	var raw_best_of string
	fmt.Scanln(&raw_best_of)
	best_of, err := strconv.Atoi(raw_best_of)
	if err != nil {
		fmt.Println("TRY AGAIN! ENTER AN INTEGER GREATER THAN 0.\nHOW MANY GAMES IN THE SERIES?")
		best_of = getGameNum()
	}

	if best_of < 1 {
		fmt.Println("TRY AGAIN! ENTER AN INTEGER GREATER THAN 0.\nHOW MANY GAMES IN THE SERIES?")
		best_of = getGameNum()
	}
	return best_of
}

func getChoice() string {
	options := map[string]bool{
		"ROCK":     true,
		"PAPER":    true,
		"SCISSORS": true,
	}
	var choice string
	fmt.Scanln(&choice)
	choice = strings.ToUpper(choice)
	if !options[choice] {
		fmt.Println("ENTER 'ROCK' 'PAPER' OR 'SCISSORS'")
		choice = getChoice()
	}
	return choice
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// ask for number of games in the series
	fmt.Println("HOW MANY ROUNDS IN THE SERIES?")
	best_of := getGameNum()

	// track game results
	wins := 0
	losses := 0
	ties := 0

	//game map gives the value against which the user choice loses (user choice = key, loses to value)
	game_map := map[string]string{
		"ROCK":     "PAPER",
		"PAPER":    "SCISSORS",
		"SCISSORS": "ROCK",
	}
	// list of options
	options := []string{
		"ROCK", "PAPER", "SCISSORS",
	}

	for i := 0; i < best_of; i++ {

		// ask for user input
		fmt.Println("YOUR TURN! MAKE A CHOICE:")

		choice := getChoice()

		// make computer choice
		rand_int := rand.Intn(3)
		comp_choice := options[rand_int]
		fmt.Println(comp_choice)

		// determine who wins the round
		if choice == comp_choice {
			fmt.Println("TIE")
			ties += 1
		} else if game_map[choice] == comp_choice {
			fmt.Println("YOU LOSE")
			losses += 1
		} else {
			fmt.Println("YOU WIN")
			wins += 1
		}
		fmt.Println(wins, ties, losses)

		if wins > (best_of / 2) {
			break
		}
		if losses > (best_of / 2) {
			break
		}
	}
	// determine series outcome
	if wins > losses {
		fmt.Println("***CONGRATS! YOU'VE WON THE SERIES!***")
	} else if wins < losses {
		fmt.Println("***SORRY! YOU'VE LOST THE SERIES!***")
	} else {
		fmt.Println("***THE SERIES HAS ENDED IN A TIE!***")
	}

}
