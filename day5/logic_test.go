package main

import (
	"testing"
)

func getTestInputs() []string {
	return []string {
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
}

func TestParseInput(t *testing.T){
	testInput := getTestInputs()[0]

	got := *parseInput(testInput)

	expected := LineSegment{startPoint: Point{xPos: 0, yPos: 9}, endPoint: Point{xPos: 5, yPos: 9}}
	if got != expected {
		t.Errorf("Testing parseInput got(%o) and expected(%o)", got, expected)
	}
}

func TestMarkIntersections(t *testing.T){
	g := new(Graph)
	g.buildGraph()

	ls := LineSegment{startPoint: Point{xPos: 1, yPos: 2}, endPoint: Point{xPos: 1, yPos: 3}}
	g.markIntersects(ls, false)

	expected := 1
	if g.points[2][1].intersections != expected {
		t.Errorf("Testing markInteresection for [1][2] got(%o) and expected(%o)", g.points[2][1].intersections, expected)
	}

	if g.points[3][1].intersections != expected {
		t.Errorf("Testing markInteresection for [1][3] got(%o) and expected(%o)", g.points[3][1].intersections, expected)
	}

	if g.points[2][2].intersections != 0 {
		t.Errorf("Testing markInteresection for [2][2] got(%o) and expected(%o)", g.points[2][2].intersections, 0)
	}
}

func TestNumPointsWithMultipleIntersects(t *testing.T) {
	g := new(Graph)
	g.buildGraph()

	inputs := getTestInputs()

	for _, i := range inputs {
		g.markIntersects(*parseInput(i), false)
	}

	got := g.numPointsWithMultipleIntersects()
	expected := 5

	if got != expected {
		t.Errorf("Testing numPointsWithMultipleIntersects got(%d) and expected(%d)", got, expected)
	}
}


func TestNumPointsWithMultipleIntersectsWithDiag(t *testing.T) {
	g := new(Graph)
	g.buildGraph()

	inputs := getTestInputs()

	for _, i := range inputs {
		g.markIntersects(*parseInput(i), true)
	}

	got := g.numPointsWithMultipleIntersects()
	expected := 12

	if got != expected {
		t.Errorf("Testing numPointsWithMultipleIntersects got(%d) and expected(%d)", got, expected)
	}
}