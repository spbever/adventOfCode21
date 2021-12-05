package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	drawnNumbers []int
	board []int
}

type BingoGame struct {
	drawnNumbers []int
	BingoBoards []BingoBoard
}

func (x BingoGame) setupGame(filePath string) BingoGame {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentRow := scanner.Text()
		if x.drawnNumbers == nil || len(x.drawnNumbers) == 0 {
			stringNums := strings.Split(currentRow, ",")

			x.drawnNumbers = []int {}
			for _, n := range stringNums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				x.drawnNumbers = append(x.drawnNumbers, num)
			}

		} else if currentRow == "" {
			if x.BingoBoards == nil {
				x.BingoBoards = [] BingoBoard {}
			}
			x.BingoBoards = append(x.BingoBoards, BingoBoard{})
		} else {
			currentBoard := x.BingoBoards[len(x.BingoBoards) - 1]
			stringNums := strings.Fields(currentRow)
			for _, n := range stringNums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				currentBoard.board = append(currentBoard.board, num)
				x.BingoBoards[len(x.BingoBoards) - 1] = currentBoard
			}
		}
	}
	return x
}

func (x BingoGame) playGame() int {
	for _, i := range x.drawnNumbers {
		for idx := 0; idx < len(x.BingoBoards); idx ++ {
			current_board := x.BingoBoards[idx]
			current_board.drawnNumbers = current_board.drawNumber(i)
			x.BingoBoards[idx] = current_board

			if current_board.checkBingo() {
				return current_board.getScore()
			}
		}
	}

	return 0
}

func (x BingoGame) playGameTillLast() int {
	for _, i := range x.drawnNumbers {
		boardsToRemove := []int {}

		for idx := 0; idx < len(x.BingoBoards); idx ++ {
			currentBoard := x.BingoBoards[idx]
			currentBoard.drawnNumbers = currentBoard.drawNumber(i)
			x.BingoBoards[idx] = currentBoard

			if currentBoard.checkBingo() {
				boardsToRemove = append(boardsToRemove, idx)
			}
		}

		if len(x.BingoBoards) == 1 {
			if x.BingoBoards[0].checkBingo() {
				return x.BingoBoards[0].getScore()
			}
		} else {
			idxFix := 0
			for _, idx := range boardsToRemove {
				x.BingoBoards = append(x.BingoBoards[:idx-idxFix], x.BingoBoards[idx-idxFix+1:]...)
				idxFix += 1
			}
		}
	}

	return 0
}

func (x BingoBoard) drawNumber(number int) []int {
	x.drawnNumbers = append(x.drawnNumbers, number)
	return x.drawnNumbers
}

func(x BingoBoard) checkBingo() bool {
	rows := []int { 0,0,0,0,0 }
	columns := []int { 0,0,0,0,0 }

	for i := 0; i < 25; i++ {
		boardNum := x.board[i]
		if contains(x.drawnNumbers, boardNum) {
			row := i / 5
			rows[row] += 1

			column := i % 5
			columns[column] += 1

			if rows[row] == 5 || columns[column] == 5 {
				return true
			}
		}
	}

	return false
}

func(x BingoBoard) getScore() int {
	score := 0

	if len(x.drawnNumbers) == 0 {
		return score
	}

	for _, i := range x.board {
		if !contains(x.drawnNumbers, i) {
			score += i
		}
	}

	return score * x.drawnNumbers[len(x.drawnNumbers)-1]
}


func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
