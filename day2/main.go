package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if !test1() || !test2() {
		return
	}

	answer1 := part1("./input.txt")
	answer2 := part2("./input.txt")

	fmt.Println("Part 2: The final location key is: " + strconv.Itoa(answer1))
	fmt.Println("Part 2: The final location key is: " + strconv.Itoa(answer2))
}

func part1(commandFile string) int {
	file, err := os.Open(commandFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	depth, horizontal := 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commandSpeed := strings.Fields(scanner.Text())
		speed, parseErr :=	strconv.Atoi(commandSpeed[1])

		if parseErr != nil {
			continue
		}

		switch commandSpeed[0] {
		case "forward":
			horizontal += speed
		case "down":
			depth += speed
		case "up":
			depth -= speed
		}
	}

	return depth * horizontal
}

func part2(commandFile string) int {
	file, err := os.Open(commandFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	depth, horizontal, aim := 0, 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commandSpeed := strings.Fields(scanner.Text())
		speed, parseErr :=	strconv.Atoi(commandSpeed[1])

		if parseErr != nil {
			continue
		}

		switch commandSpeed[0] {
		case "forward":
			horizontal += speed
			depth += aim * speed
		case "down":
			aim += speed
		case "up":
			aim -= speed
		}
	}

	return depth * horizontal
}

func test1() bool {
	answer := part1("./test_input.txt")
	if answer != 150 {
		fmt.Println("Part1 Test = " + strconv.Itoa(answer) + "; wanted 150")
		return false
	}
	return true
}

func test2() bool {
	answer := part2("./test_input.txt")
	if answer != 900 {
		fmt.Println("Part2 Test = " + strconv.Itoa(answer) + "; wanted 900")
		return false
	}
	return true
}
