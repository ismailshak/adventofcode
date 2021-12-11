// Advent of Code Day 9: Smoke Basin
// incomplete, no longer have bandwidth, come back to this one later and submit part 2.
package day09

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Smoke Basin")
	fmt.Println("-------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineRiskLevels(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineBasins(inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

func parseRow(value string) []int {
	slice := []int{}
	for _, char := range value {
		slice = append(slice, util.ParseInt(string(char)))
	}

	return slice
}

func isEmpty(value string) bool {
	return value == ""
}

func isHeightLowest(height, up, down, left, right int) bool {
	return height < up && height < down && height < left && height < right
}

func getHeightAtIndex_String(i int, value string) int {
	// return the max if we're out of the grid's bounds
	if isEmpty(value) || i >= len(value) {
		return 9
	}

	return util.ParseInt(string(value[i]))
}

func getHeightAtIndex_Slice(i int, value []int) int {
	// return the max if we're out of the grid's bounds
	if i < 0 || i >= len(value) {
		return 9
	}

	return value[i]
}

func findLowestPoints(row []int, next, previous *bufio.Scanner) []int {
	previousInput := previous.Text()
	nextInput := next.Text()
	lowestPointsFound := []int{}

	for i, height := range row {
		up := getHeightAtIndex_String(i, previousInput)
		down := getHeightAtIndex_String(i, nextInput)
		left := getHeightAtIndex_Slice(i-1, row)
		right := getHeightAtIndex_Slice(i+1, row)

		if isHeightLowest(height, up, down, left, right) {
			lowestPointsFound = append(lowestPointsFound, height)
		}
	}

	return lowestPointsFound
}

func heightToRiskLevel(height int) int {
	return height + 1
}

func sum(list []int) int {
	sum := 0
	for _, height := range list {
		sum += heightToRiskLevel(height)
	}
	return sum
}

func determineRiskLevels(inputFileName string) int {
	// creating duplicate file buffers so I can spawn separate scanners, that read from
	// different positions simultaneously
	files := [3]*os.File{
		util.OpenInputFile(9, inputFileName),
		util.OpenInputFile(9, inputFileName),
		util.OpenInputFile(9, inputFileName),
	}
	defer files[0].Close()
	defer files[1].Close()
	defer files[2].Close()

	// to start, current and previous will be at the same line (1) and next will be on line 2
	previous := bufio.NewScanner(files[0])
	current := bufio.NewScanner(files[1])
	next := bufio.NewScanner(files[2])
	next.Scan()

	lowestPoints := []int{}

	for current.Scan() {
		if !isEmpty(next.Text()) {
			next.Scan()
		}

		row := parseRow(current.Text())
		lowestPoints = append(lowestPoints, findLowestPoints(row, next, previous)...)

		previous.Scan()
	}

	return sum(lowestPoints)
}

func createGrid(file []byte) [][]int {
	input := strings.Split(string(file), "\n")
	grid := make([][]int, 100)

	for i, line := range input {
		points := strings.Split(line, "")
		grid[i] = make([]int, 100)
		for j, char := range points {
			grid[i][j] = util.ParseInt(char)
		}
	}

	return grid
}

//func countPointsInBasin(grid [][]int, i, j int) int {
//	if grid[i][j] != 9 {

//	}

//	if i != 100 {
//		return countPointsInBasin()
//	}
//	return 0
//}

// lazy, tired, swallowing the whole input into memory...
func determineBasins(inputFileName string) int {
	file, err := os.ReadFile(util.BuildPuzzlePath(9, inputFileName))
	if err != nil {
		panic(err)
	}

	grid := createGrid(file)

	for _, row := range grid {
		for _, point := range row {
			if point == 9 {
				continue
			}
			//countPointsInBasin(grid, i, j)
		}
	}

	return 0
}
