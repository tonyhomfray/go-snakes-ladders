package snakelib

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatePlayers(numberOfPlayers *int) []player {

	fmt.Print("How many people are playing? ")
	fmt.Scanln(numberOfPlayers)

	// fmt.Printf("%v\n", *numberOfPlayers)

	players := make([]player, *numberOfPlayers)

	for i := 0; i < *numberOfPlayers; i++ {
		playerX := CreatePlayer(i)
		players[i] = playerX
	}
	fmt.Println()
	Pause(1000)
	fmt.Print("Let's play!\n\n")

	return players
}

func CreatePlayer(playerNumber int) player {
	var name string
	fmt.Printf("Enter the name of Player %v: ", playerNumber+1)
	fmt.Scan(&name)

	playerX := player{name: name}
	playerX.currentSquare = 0

	return playerX
}

func RollDice() int {
	return (rand.Intn(6) + 1)
}

func CheckForSnakes(landedOn int) bool {
	if _, snake := snakes[landedOn]; snake {
		return true
	} else {
		return false
	}
}

func CheckForLadders(landedOn int) bool {

	if _, ladder := ladders[landedOn]; ladder {
		return true
	} else {
		return false
	}
}

func CheckIfWon(square *int) bool {
	if *square >= endSquare {
		*square = endSquare
		fmt.Println("We have a winner!!!")
		weHaveAWinner = true
		return true
	}
	return false
}

func ShowSnake() {
	fmt.Println(
		`
	 ____
	/ o o\
	\  ---<
	 \  /
_________/ /
-=:___________/`)
}

func ShowLadder() {
	fmt.Println(
		`
   H
   H
   H
   H
   H`)
}

func GetNextPlayer(currentPlayer *int, numberOfPlayers int) int {
	if *currentPlayer+1 == numberOfPlayers {
		return 0
	} else {
		return *currentPlayer + 1
	}
}

func EndGame(winner string) {
	fmt.Printf("\n%v WWWWWIIIINNNNNSSSSS!!!!!!!\n\n", winner)
}

func ThanksForPlaying() {
	fmt.Println()
	fmt.Println("Thanks for playing!")
	fmt.Println()
}

func Pause(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
