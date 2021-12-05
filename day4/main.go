package main

import "fmt"

func main() {
	bingoGame := new(BingoGame)
	updatedGame := bingoGame.setupGame("./input.txt")
	finalScore := updatedGame.playGame()
	lastToWinScore := updatedGame.playGameTillLast()

	fmt.Printf("Part1: The score of the winning board is %d\n", finalScore)
	fmt.Printf("Part2: The score of the board last to win is %d\n", lastToWinScore)
}