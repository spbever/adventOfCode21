package main

import (
	"testing"
)

func getTestBoard() []int {
	var b []int

	for i := 1; i <= 25; i++ {
		b = append(b, i)
	}

	return b
}

func arrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSetupGameCheckDrawnNumbers (t *testing.T) {
	bingoGame := new(BingoGame)
	updatedGame := bingoGame.setupGame("./test_input.txt")

	expected := []int {7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}

	got := updatedGame.drawnNumbers

	if !arrayEqual(got, expected) {
		t.Errorf("Testing playGame got(%d) and expected(%d)", got, expected)
	}
}

func TestPlayGame(t *testing.T) {
	bingoGame := new(BingoGame)
	updatedGame := bingoGame.setupGame("./test_input.txt")

	got := updatedGame.playGame()
	expected := 4512
	if got != expected {
		t.Errorf("Testing playGame got(%d) and expected(%d)", got, expected)
	}
}

func TestPlayGameTillLast(t *testing.T) {
	bingoGame := new(BingoGame)
	updatedGame := bingoGame.setupGame("./test_input.txt")

	got := updatedGame.playGameTillLast()
	expected := 1924
	if got != expected {
		t.Errorf("Testing playGame got(%d) and expected(%d)", got, expected)
	}
}

func TestGetScoreWithNoDrawnNumbers (t *testing.T) {
	bingoBoard := new(BingoBoard)

	expected := 0
	got := bingoBoard.getScore()

	if got != expected {
		t.Errorf("Testing getScore with no drawn numbers got(%d) and expected(%d)", got, expected)
	}
}


func TestGetScoreWithBingoRow (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,2,3,4,5}
	bingoBoard.board = getTestBoard()
    expected := 1550

	got := bingoBoard.getScore()

	if got != expected {
		t.Errorf("Testing getScore with bingo got(%d) and expected(%d)", got, expected)
	}
}

func TestCheckBingForRow (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,2,3,4,5}
	bingoBoard.board = getTestBoard()
	expected := true

	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing checkBingForRow with row bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestCheckBingForRowFalse (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,2,3,4,6}
	bingoBoard.board = getTestBoard()
	expected := false

	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing checkBingForRow with no row bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestCheckBingForColumn (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,6,11,16,21}
	bingoBoard.board = getTestBoard()
	expected := true

	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing checkBingForColumn with column bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestCheckBingForColumnFalse (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,6,11,16,20}
	bingoBoard.board = getTestBoard()
	expected := false

	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing checkBingForColumn with no column bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestGetScoreWithBingoColumn (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,6,11,16,21}
	bingoBoard.board = getTestBoard()
	expected := 5670

	got := bingoBoard.getScore()

	if got != expected {
		t.Errorf("Testing getScore with bingo got(%d) and expected(%d)", got, expected)
	}
}

func TestDrawNumberWithoutBingo (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,6,11}
	bingoBoard.board = getTestBoard()
	expected := false
	bingoBoard.drawnNumbers = bingoBoard.drawNumber(16)
	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing DrawNumberWithoutBingo with out bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestDrawNumberWithBingo (t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.drawnNumbers = []int {1,6,11,21}
	bingoBoard.board = getTestBoard()
	expected := true
	bingoBoard.drawnNumbers = bingoBoard.drawNumber(16)
	got := bingoBoard.checkBingo()

	if got != expected {
		t.Errorf("Testing DrawNumberWithBingo with bingo got(%t) and expected(%t)", got, expected)
	}
}

func TestDrawNumberAddsNumberToDrawnNumbers(t *testing.T) {
	bingoBoard := new(BingoBoard)
	bingoBoard.board = getTestBoard()
	expected := []int {1,2}

	bingoBoard.drawnNumbers = bingoBoard.drawNumber(1)
	bingoBoard.drawnNumbers = bingoBoard.drawNumber(2)

	got := bingoBoard.drawnNumbers

	if !arrayEqual(got, expected) {
		t.Errorf("Testing DrawNumberWithBingo with bingo got(%d) and expected(%d)", got, expected)
	}
}
