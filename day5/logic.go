package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Graph struct {
	points [1000][1000]Point
}

func(g *Graph) buildGraph() {
	for x := 0; x <= 999; x ++ {
		for y := 0; y <= 999; y ++ {
			p := new(Point)
			p.buildPoint(x, y)
			g.points[x][y] = *p
		}
	}
}

func(g *Graph) markIntersects(ls LineSegment, withDiagnols bool){
	xMin := ls.startPoint.xPos
	xMax := ls.endPoint.xPos

	yMin := ls.startPoint.yPos
	yMax := ls.endPoint.yPos


	if yMin != yMax && xMin != xMax {
		if withDiagnols {

			xModifier := 1
			if xMin > xMax {
				xModifier = -1
			}
			yModifier := 1
			if yMin > yMax {
				yModifier = -1
			}

			//Assumed 45 degree angle or a slope of 1/1
			xCurrent := xMin
			for yCurrent := yMin; yCurrent != yMax; yCurrent += yModifier {
				g.points[yCurrent][xCurrent].intersections += 1
				xCurrent += xModifier
			}
			g.points[yMax][xMax].intersections += 1
		}
		return
	}

	if yMin > yMax {
		t := yMax
		yMax = yMin
		yMin = t
	}

	if xMin > xMax {
		t := xMax
		xMax = xMin
		xMin = t
	}

	for x := xMin; x <= xMax; x ++ {
		for y := yMin; y <= yMax; y ++ {
			g.points[y][x].intersections += 1
		}
	}
}

func(g *Graph) numPointsWithMultipleIntersects() int{
	total := 0
	for x := 0; x <= 999; x ++ {
		for y := 0; y <= 999; y ++ {
			if g.points[y][x].intersections > 1 {
				total += 1
			}
		}
	}
	return total
}

func(g *Graph) printGraph() {
	for _, a := range g.points {
		for _, b := range a {
			fmt.Printf("%d", b.intersections)
		}
		fmt.Printf("\n")
	}
}

type Point struct {
	xPos int
	yPos int
	intersections int
}

func(pt *Point) buildPoint(x int, y int) {
	pt.xPos = x
	pt.yPos = y
	pt.intersections = 0
}

type LineSegment struct {
	startPoint Point
	endPoint Point
}

func (ls *LineSegment) buildLineSegment(sPoint Point, ePoint Point) {
	ls.startPoint = sPoint
	ls.endPoint = ePoint
}

func parseInput(input string) *LineSegment {
	parts := strings.Split(input, " -> ")
	inputs1 := strings.Split(parts[0], ",")
	var point1 = new(Point)
	point1.buildPoint(parseIntFromString(inputs1[0]), parseIntFromString(inputs1[1]))

	inputs2 := strings.Split(parts[1], ",")
	var point2 = new(Point)
	point2.buildPoint(parseIntFromString(inputs2[0]), parseIntFromString(inputs2[1]))

	var lineSegment = new(LineSegment)

	lineSegment.buildLineSegment(*point1, *point2)

	return lineSegment
}

func parseIntFromString(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}