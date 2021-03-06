// Advent of Code Day 5: Hydrothermal Venture
package day05

import (
	"aoc/util"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Hydrothermal Venture")
	fmt.Println("----------------------------")

	// Part 1
	p1Start := time.Now()
	noOfPoints := determinePoints4Axis(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", noOfPoints, p1Duration)

	// Part 2
	p2Start := time.Now()
	noOfPoints = determinePoints8Axis(inputFileName) // with diagonals
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", noOfPoints, p2Duration)

	// Fun visualization
	drawVents(util.BuildPuzzlePath(5, "vents.png"), inputFileName, 1000, 1000)

	// Part 3?? (see README)
	drawVents(util.BuildPuzzlePath(5, "reddit-bonus.png"), "reddit-input.txt", 2000, 2000)
}

type Orientation uint8

const (
	HORIZONTAL Orientation = 0
	VERTICAL   Orientation = 1
	DIAGONAL   Orientation = 2
)

type Point struct {
	x int
	y int
}

func (point *Point) ToString() string {
	return strconv.Itoa(int(point.x)) + "," + strconv.Itoa(int(point.y))
}

func parsePoint(point string) (int, int) {
	pointParts := strings.Split(point, ",")
	return util.ParseInt(pointParts[0]), util.ParseInt(pointParts[1])
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

func direction(orientation Orientation, start, end Point) int {
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

func diagonalDirection(i, j int, start, end Point) (int, int) {
	var nextX, nextY int
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

func withinRange(index int, orientation Orientation, start, end Point) bool {
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
func diagonalWithinRange(i, j int, start, end Point) bool {
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

func determinePoints(considerDiagonal bool, inputFileName string) int {
	inputFile := util.OpenInputFile(5, inputFileName)
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
func determinePoints4Axis(inputFileName string) int {
	return determinePoints(false, inputFileName)
}

// i.e. horizontal + vertical + diagonal
func determinePoints8Axis(inputFileName string) int {
	return determinePoints(true, inputFileName)
}

//=========================================

func drawLine(canvas *image.RGBA, start, end Point) {
	pixelColor := color.RGBA{R: 233, G: 166, B: 166, A: 255}
	if isVentHorizontal(start, end) {
		for i := start.x; withinRange(i, HORIZONTAL, start, end); i += direction(HORIZONTAL, start, end) {
			canvas.SetRGBA(int(i), int(start.y), pixelColor)
		}
	}

	if isVentVertical(start, end) {
		for j := start.y; withinRange(j, VERTICAL, start, end); j += direction(VERTICAL, start, end) {
			canvas.SetRGBA(int(start.x), int(j), pixelColor)
		}
	}

	if isVentDiagonal(start, end) {
		for i, j := start.x, start.y; diagonalWithinRange(i, j, start, end); i, j = diagonalDirection(i, j, start, end) {
			canvas.SetRGBA(int(i), int(j), pixelColor)
		}
	}
}

func drawVents(fileName, inputName string, width, length int) {
	inputFile := util.OpenInputFile(5, inputName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	canvas := image.NewRGBA(image.Rect(0, 0, width, length))

	for input.Scan() {
		start, end := parseLine(input.Text())
		drawLine(canvas, start, end)
	}

	imageFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = png.Encode(imageFile, canvas)
	if err != nil {
		panic(err)
	}
	fmt.Println("Image drawn and saved to", fileName)
}
