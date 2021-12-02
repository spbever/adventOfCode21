package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1 () {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	largerCount := 0
	var previous int64
	first := true

	for scanner.Scan() {
		current, parseErr := strconv.ParseInt(scanner.Text(), 10, 64)
		if parseErr != nil {
			log.Fatal(parseErr)
			continue
		}
		if first {
			previous = current
			first = false
			continue
		}
		if current > previous {
			largerCount += 1
		}
		previous = current
	}

	fmt.Println("Part 1: There are " + strconv.Itoa(largerCount) + " numbers larger than the previous." )
}

func part2 () {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Lets assume the water has to be greater depth than 0
	scanner := bufio.NewScanner(file)
	largerCount := 0
	previous1, previous2, previous3  := 0,0,0

	for scanner.Scan() {
		current, parseErr := strconv.Atoi(scanner.Text())
		if parseErr != nil {
			log.Fatal(parseErr)
			continue
		}
		if previous1 == 0 {
			previous1 = current
			continue
		}

		if previous2 == 0 {
			previous2 = current
			continue
		}

		if previous3 == 0 {
			previous3 = current
			continue
		}

		if previous2 + previous3 + current > previous1 + previous2 + previous3 {
			largerCount +=1
		}

		previous1 = previous2
		previous2 = previous3
		previous3 = current
	}
	fmt.Println("Part 2: There are " + strconv.Itoa(largerCount) + " number sets larger than the previous set." )
}