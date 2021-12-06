package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	g := new(Graph)
	g.buildGraph()
	g2 := new(Graph)
	g2.buildGraph()

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		g.markIntersects(*parseInput(scanner.Text()), false)
		g2.markIntersects(*parseInput(scanner.Text()), true)
	}

	fmt.Printf("Part1: Number of points with multiple intersections: %d\n", g.numPointsWithMultipleIntersects())
	fmt.Printf("Part2: Number of points with multiple intersections: %d\n", g2.numPointsWithMultipleIntersects())

}
