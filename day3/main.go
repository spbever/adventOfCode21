package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bitlist, _ := readLines("./input.txt")
	fmt.Printf("Part 1: The power consumption is: %d\n", calcPoserConsumption(bitlist))
	fmt.Printf("Part 2: The life support rating is: %d\n", calcLifeSupportRating(bitlist))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
