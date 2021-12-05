// Advent of Code Day 5: Hydrothermal Venture
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\nSolving Hydrothermal Venture")
	fmt.Println("----------------------------")

	// Part 1
	p1Start := time.Now()
	noOfPoints := determinePoints4Axis()
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", noOfPoints, p1Duration)

	// Part 2
	p2Start := time.Now()
	noOfPoints = determinePoints8Axis() // with diagonals
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", noOfPoints, p2Duration)
}

type Orientation uint8

const (
	HORIZONTAL Orientation = 0
	VERTICAL   Orientation = 1
	DIAGONAL   Orientation = 2
)

type Point struct {
	x int32
	y int32
}

func (point *Point) ToString() string {
	return strconv.Itoa(int(point.x)) + "," + strconv.Itoa(int(point.y))
}

func openInputFile() *os.File {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic("Error. Failed to read input file.")
	}
	return inputFile
}

func parseInt(value string) int32 {
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic("Error. Failed to parse int")
	}
	return int32(intValue)
}

func parsePoint(point string) (int32, int32) {
	pointParts := strings.Split(point, ",")
	return parseInt(pointParts[0]), parseInt(pointParts[1])
}

func parseLine(line string) (Point, Point) {
	lineParts := strings.Fields(line)
	x1, y1 := parsePoint(lineParts[0])
	x2, y2 := parsePoint(lineParts[2])
	return Point{x: x1, y: y1}, Point{x: x2, y: y2}
}

func isVentHorizontal(start, end Point) bool {
	return start.y == end.y
}

func isVentVertical(start, end Point) bool {
	return start.x == end.x
}

// for readability's sake
func isVentDiagonal(start, end Point) bool {
	return !isVentHorizontal(start, end) && !isVentVertical(start, end)
}

func direction(orientation Orientation, start, end Point) int32 {
	switch orientation {
	case HORIZONTAL:
		if start.x > end.x {
			return -1
		}
		return 1
	case VERTICAL:
		if start.y > end.y {
			return -1
		}
		return 1
	}
	panic("Error. Direction didn't match HOR/VER")
}

func diagonalDirection(i, j int32, start, end Point) (int32, int32) {
	var nextX, nextY int32
	if start.x > end.x {
		nextX = -1
	} else {
		nextX = 1
	}

	if start.y > end.y {
		nextY = -1
	} else {
		nextY = 1
	}

	return i + nextX, j + nextY
}

func withinRange(index int32, orientation Orientation, start, end Point) bool {
	switch orientation {
	case HORIZONTAL:
		if start.x > end.x {
			return index >= end.x
		}
		return index <= end.x
	case VERTICAL:
		if start.y > end.y {
			return index >= end.y
		}
		return index <= end.y
	}

	panic("Error. withinRange didn't get HOR/VER")
}

// getting tired/lazy - just go with it
func diagonalWithinRange(i, j int32, start, end Point) bool {
	var xCondition, yCondition bool
	xDirection, yDirection := diagonalDirection(0, 0, start, end)

	if xDirection > 0 {
		xCondition = i <= end.x
	} else {
		xCondition = i >= end.x
	}

	if yDirection > 0 {
		yCondition = j <= end.y
	} else {
		yCondition = j >= end.y
	}

	return xCondition && yCondition
}

func addPoint(pointFrequency *map[string]int, multiVentPoints *map[string]bool, newPoint, start Point) {
	(*pointFrequency)[newPoint.ToString()] += 1
	// if the point has been visited more than once, and is not yet recorded as multi
	if (*pointFrequency)[newPoint.ToString()] > 1 && !(*multiVentPoints)[newPoint.ToString()] {
		(*multiVentPoints)[newPoint.ToString()] = true
	}
}

// idk why im doing it like this but im in too deep now (pun not intended)
func buildSegment(pointFrequency *map[string]int, multiVentPoints *map[string]bool, start, end Point, considerDiagonal bool) {
	if isVentHorizontal(start, end) {
		for i := start.x; withinRange(i, HORIZONTAL, start, end); i += direction(HORIZONTAL, start, end) {
			addPoint(pointFrequency, multiVentPoints, Point{x: i, y: start.y}, start)
		}
	}

	if isVentVertical(start, end) {
		for i := start.y; withinRange(i, VERTICAL, start, end); i += direction(VERTICAL, start, end) {
			addPoint(pointFrequency, multiVentPoints, Point{x: start.x, y: i}, start)
		}
	}

	// Diagonals assumed to be always at 45 degrees
	if considerDiagonal && isVentDiagonal(start, end) {
		for i, j := start.x, start.y; diagonalWithinRange(i, j, start, end); i, j = diagonalDirection(i, j, start, end) {
			addPoint(pointFrequency, multiVentPoints, Point{x: i, y: j}, start)
		}
	}
}

func determinePoints(considerDiagonal bool) int {
	inputFile := openInputFile()
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	pointFrequency := make(map[string]int)
	multiVentPoints := make(map[string]bool)
	for input.Scan() {
		start, end := parseLine(input.Text())
		buildSegment(&pointFrequency, &multiVentPoints, start, end, considerDiagonal)
	}
	return len(multiVentPoints)
}

// i.e. horizontal + vertical
func determinePoints4Axis() int {
	return determinePoints(false)
}

// i.e. horizontal + vertical + diagonal
func determinePoints8Axis() int {
	return determinePoints(true)
}
