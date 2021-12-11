// Advent of Code Day 11: Dumbo Octopus
package day11

import (
	"aoc/util"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Dumbo Octopus")
	fmt.Println("---------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineFlashes(100, inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := findStepWithMaxFlashes(1000, inputFileName) // idk some super large number
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

func printGrid(grid *[10][10]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf(" %3d", col)
		}
		fmt.Print("\n")
	}
	fmt.Print("---------\n")
}

func bumpEnergyLevels(grid *[10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j]++
		}
	}
}

// terrible.
func simulateFlashes(grid *[10][10]int) int {
	flashCount := 0
	didIterationFlash := true

	for didIterationFlash {
		didIterationFlash = false
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j] > 9 {
					grid[i][j] = 0
					//printGrid(grid)
					flashCount++
					didIterationFlash = true
					up := (i - 1) >= 0
					down := (i + 1) < 10
					left := (j - 1) >= 0
					right := (j + 1) < 10

					if up && grid[i-1][j] != 0 {
						grid[i-1][j] += 1
					}
					if down && grid[i+1][j] != 0 {
						grid[i+1][j] += 1
					}
					if left && grid[i][j-1] != 0 {
						grid[i][j-1] += 1
					}
					if right && grid[i][j+1] != 0 {
						grid[i][j+1] += 1
					}
					if up && left && grid[i-1][j-1] != 0 {
						grid[i-1][j-1] += 1
					}
					if up && right && grid[i-1][j+1] != 0 {
						grid[i-1][j+1] += 1
					}
					if down && left && grid[i+1][j-1] != 0 {
						grid[i+1][j-1] += 1
					}
					if down && right && grid[i+1][j+1] != 0 {
						grid[i+1][j+1] += 1
					}
				}
			}
		}
	}

	return flashCount
}

func determineFlashes(steps int, inputFileName string) int {
	inputFile := util.OpenInputFile(11, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	grid := [10][10]int{}

	// build grid
	for i := 0; i < 10; i++ {
		input.Scan()
		line := strings.Split(input.Text(), "")
		for j := 0; j < 10; j++ {
			grid[i][j] = util.ParseInt(line[j])
		}
	}

	// simulate evolution
	totalFlashes := 0
	for k := 0; k < steps; k++ {
		bumpEnergyLevels(&grid)
		flashCount := simulateFlashes(&grid)
		// for part 2, I'll pass a super large steps value so that when
		// all octopi flash the return will be the iteration (vs the flash count :shrug:)
		if flashCount == 100 {
			return k
		}

		totalFlashes += flashCount
	}

	return totalFlashes
}

func findStepWithMaxFlashes(steps int, inputFileName string) int {
	return determineFlashes(steps, inputFileName) + 1 // adding 1 because part 1's function started at step 0
}
