package main

import (
	"fmt"
	sl "snakes/snakelib"
)

type player struct {
	name          string
	currentSquare int
}

var snakes = map[int]int{16: 6, 47: 26, 49: 11, 56: 53, 62: 18, 64: 60, 87: 24, 92: 73, 95: 75, 98: 78}
var ladders = map[int]int{2: 38, 4: 14, 8: 31, 21: 42, 36: 44, 51: 67, 71: 91, 80: 100}

const endSquare int = 25

var numberOfPlayers int
var weHaveAWinner bool = false

func main() {
	fmt.Println()
	fmt.Println("##############################")
	fmt.Println("Welcome to Snakes and Ladders!")
	fmt.Println("##############################")
	fmt.Println()
	sl.Pause(2000)
	players := sl.CreatePlayers(&numberOfPlayers)
	currentPlayer := 0

	for !weHaveAWinner {
		fmt.Printf("%v's turn\n\n", players[currentPlayer].name)
		sl.Pause(600)

		fmt.Println("Rolling")
		sl.Pause(400)
		roll := sl.RollDice()

		for i := 0; i < roll; i++ {
			fmt.Print("# ")
			sl.Pause(400)
		}
		fmt.Println()
		sl.Pause(600)
		fmt.Printf("%v rolled a %v ... \n", players[currentPlayer].name, roll)
		sl.Pause(700)
		landedOn := players[currentPlayer].currentSquare + roll
		sl.Pause(500)
		fmt.Printf("... and landed on %v\n\n", landedOn)
		sl.Pause(500)

		if sl.CheckForSnakes(landedOn) {
			landedOn = snakes[landedOn]
			fmt.Printf("!!! SNAKE !!!\n")
			sl.ShowSnake()

			fmt.Printf("\nGo back to %v!\n\n", landedOn)

		}

		if sl.CheckForLadders(landedOn) {
			landedOn = ladders[landedOn]
			fmt.Printf("!!! LADDER !!!\n")
			sl.ShowLadder()
			fmt.Printf("\nGo forward to %v!\n\n", landedOn)

		}

		if sl.CheckIfWon(&landedOn) {
			sl.EndGame(players[currentPlayer].name)
			break
		}

		players[currentPlayer].currentSquare = landedOn

		currentPlayer = sl.GetNextPlayer(&currentPlayer, numberOfPlayers)
		sl.Pause(2000)
	}

	sl.ThanksForPlaying()

}
