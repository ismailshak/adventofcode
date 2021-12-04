// Advent of Code Day 4: Giant Squid.
// Good one to come back to later, was the first time I realized slices are just pointers
// to an underlying array (with a little bit of extra metadata), so passing a pointer to a slice
// as a func arg is literally useless. Clever and weird construct for dynamically sized arrays, ngl.
// TODO: Read up on slices => come back and clean up.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	INPUT_FILE = "./input.txt"
	GRID_SIZE  = 5
	RESET      = "\033[0m"
	CYAN       = "\033[36m"
	PURPLE     = "\033[35m"
)

func main() {
	fmt.Println("\nSolving Giant Squid")
	fmt.Println("-------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineWinningScore()
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	fmt.Println()

	// Part 2
	p2Start := time.Now()
	p2Result := determineLosingScore()
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

type Cell struct {
	value  int32
	marked bool
}

type Board struct {
	grid [GRID_SIZE][GRID_SIZE]Cell
}

func (b *Board) Print() {
	fmt.Println("-------------------")
	for _, row := range b.grid {
		for _, col := range row {
			if col.marked {
				fmt.Print(PURPLE)
			}
			fmt.Printf(" %2d ", col.value)
			fmt.Print(RESET)
		}
		fmt.Print("\n")
	}
	fmt.Println("-------------------")
}

func (b *Board) MarkIfExists(drawnNumber int32) (found bool, foundRow, foundCol int) {
out:
	for i := 0; i < len(b.grid); i++ {
		for j := 0; j < len(b.grid[i]); j++ {
			if b.grid[i][j].value == drawnNumber {
				b.grid[i][j].marked = true
				found = true
				foundRow = i
				foundCol = j
				break out
			}
		}
	}
	return
}

func (b *Board) CheckRowWin(rowIndex int) bool {
	for _, cell := range b.grid[rowIndex] {
		if !cell.marked {
			return false
		}
	}
	return true
}

func (b *Board) CheckColWin(colIndex int) bool {
	for i := 0; i < GRID_SIZE; i++ {
		cell := b.grid[i][colIndex]
		if !cell.marked {
			return false
		}
	}
	return true
}

func (b *Board) Score(drawnNumber int32) int32 {
	sum := int32(0)
	for _, row := range b.grid {
		for _, col := range row {
			if !col.marked {
				sum += col.value
			}
		}
	}
	return sum * drawnNumber
}

func parseFile() []string {
	file, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		panic("Error. Failed to read file")
	}

	input := strings.Split(string(file), "\n")
	return input
}

func parseInt(value string) int32 {
	output, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic("Error. Failed to parse int")
	}
	return int32(output)
}

func buildBoard(boards []Board, boardNumber, rowCounter int32, line string) {
	cellValues := strings.Fields(line)

	for column, value := range cellValues {
		parsedValue := parseInt(string(value))
		boards[boardNumber].grid[rowCounter][column].value = parsedValue
	}
}

func createBoards(input []string) []Board {
	boards := []Board{}
	var boardNumber int32
	var rowCounter int32

	for lineNum, line := range input {
		// skip over first two lines, don't need 'em
		if lineNum == 0 || lineNum == 1 {
			continue
		}
		// if empty string, we're about to encounter a new board
		if line == "" {
			boardNumber++
			rowCounter = 0
			continue
		}
		// create a board if we haven't yet
		if len(boards) == int(boardNumber) {
			boards = append(boards, Board{})
		}
		// add data to board, line by line
		buildBoard(boards, boardNumber, rowCounter, line)
		rowCounter++
	}

	return boards
}

func prepareBingo() ([]string, []Board) {
	input := parseFile()
	draw := strings.Split(input[0], ",")
	boards := createBoards(input)

	return draw, boards
}

func determineWinningScore() int32 {
	draw, boards := prepareBingo()
	for i := 0; i < len(draw); i++ {
		for j := 0; j < len(boards); j++ {
			drawnNumber := parseInt(draw[i])
			board := &(boards)[j]
			found, row, col := board.MarkIfExists(drawnNumber)
			if found {
				completedRow := board.CheckRowWin(row)
				completedCol := board.CheckColWin(col)
				if completedCol || completedRow {
					fmt.Printf("Winning board (draw: %v)\n", drawnNumber)
					board.Print()
					return board.Score(drawnNumber)
				}
			}
		}
	}
	return 0
}

// to be a little memory and speed efficient, just stopping once every board
// has won at least once
func addToWinners(winners *map[int]bool, maxWinners, winnerIndex int) bool {
	(*winners)[winnerIndex] = true
	if len(*winners) == maxWinners {
		return true
	}
	return false
}

func determineLosingScore() int32 {
	draw, boards := prepareBingo()
	winners := make(map[int]bool)
	for i := 0; i < len(draw); i++ {
		for j := 0; j < len(boards); j++ {
			drawnNumber := parseInt(draw[i])
			board := &(boards)[j]
			found, row, col := board.MarkIfExists(drawnNumber)
			if found {
				completedRow := board.CheckRowWin(row)
				completedCol := board.CheckColWin(col)
				if completedCol || completedRow {
					lastWinner := addToWinners(&winners, len(boards), j)
					if lastWinner {
						fmt.Printf("Losing board (draw: %v)\n", drawnNumber)
						board.Print()
						return board.Score(drawnNumber)
					}
				}
			}
		}
	}
	return 0
}
