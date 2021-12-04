package main

import (
	"strconv"
)

func calcLifeSupportRating(bitList []string) int {
	return calcOxygenGeneratorRating(bitList) * calcCo2ScrubberRating(bitList)
}

func calcOxygenGeneratorRating(bitList []string) int {
	if len(bitList) == 0 {
		return 0
	}
	size := len(bitList[0])

	currentList := bitList

	for i := 1; i <= size; i++ {
		if len(currentList) == 1 {
			break
		}
		currentList = filterBitByPosition(currentList, i, mostCommon(currentList, i))
	}

	return bitStringToInt(currentList[0])
}

func calcCo2ScrubberRating(bitList []string) int {
	if len(bitList) == 0 {
		return 0
	}
	size := len(bitList[0])

	currentList := bitList

	for i := 1; i <= size; i++ {
		if len(currentList) == 1 {
			break
		}

		lc := "0"
		if mostCommon(currentList, i) == "0" {
			lc = "1"
		}

		currentList = filterBitByPosition(currentList, i, lc )
	}

	return bitStringToInt(currentList[0])
}

func filterBitByPosition(bitList []string, position int, filter string) []string {
	var newList []string

	for  _, bitString := range bitList {
		if string(bitString[position-1]) == filter {
			newList = append(newList, bitString)
		}
	}

	return newList
}

func calcPoserConsumption(bitList []string) int {
	return calcGamaRate(bitList) * calcEpsilonRate(bitList)
}

func calcGamaRate(bitList []string) int {
	if len(bitList) == 0 {
		return 0
	}
	size := len(bitList[0])
	gamaBitMap := ""

	for i := 1; i <= size; i++ {
		gamaBitMap += mostCommon(bitList, i)
	}

	return bitStringToInt(gamaBitMap)
}

func calcEpsilonRate(bitList []string) int {
	if len(bitList) == 0 {
		return 0
	}
	size := len(bitList[0])
	epsilonBitMap := ""

	for i := 1; i <= size; i++ {
		if mostCommon(bitList, i) == "1" {
			epsilonBitMap += "0"
		} else {
			epsilonBitMap += "1"
		}
	}

	return bitStringToInt(epsilonBitMap)
}


func mostCommon (bitList []string, position int) string{
	total := 0
	for _, b := range bitList {
		if b[position-1] == '1' {
			total += 1
		}
	}
	if float64(total) / float64(len(bitList)) >= 0.5 {
		return "1"
	} else {
		return "0"
	}
}

func bitStringToInt (bitString string) int {
	if i, err := strconv.ParseInt(bitString, 2, 64); err != nil {
		return -1
	} else {
		return int(i)
	}
}
