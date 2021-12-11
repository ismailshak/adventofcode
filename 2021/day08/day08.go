// Advent of Code Day 8: Seven Segment Search
package day08

import (
	"aoc/util"
	"bufio"
	"fmt"
	"math"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Seven Segment Search")
	fmt.Println("-------------------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineSimpleDigitsCount(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineOutputSum(inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

func parseInputLine(line string) ([]string, []string) {
	inputParts := strings.Fields(line)    // splits a string on whitespace, and trims each element
	delimiterIndex := len(inputParts) - 5 // we always have 4 digits after the delimiter
	return inputParts[:delimiterIndex], inputParts[delimiterIndex+1:]
}

func countOutputDigitFrequency(digitFrequency *map[int]int, outputDigits []string) {
	for _, digit := range outputDigits {
		switch len(string(digit)) {
		case 2:
			(*digitFrequency)[1] += 1
		case 3:
			(*digitFrequency)[7] += 1
		case 4:
			(*digitFrequency)[4] += 1
		case 7:
			(*digitFrequency)[8] += 1
		}
	}
}

func determineSimpleDigitsCount(inputFileName string) int {
	inputFile := util.OpenInputFile(8, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	digitFrequency := make(map[int]int)

	for input.Scan() {
		_, outputValue := parseInputLine(input.Text())
		countOutputDigitFrequency(&digitFrequency, outputValue)
	}

	return digitFrequency[1] + digitFrequency[4] + digitFrequency[7] + digitFrequency[8]
}

// 1, 4, 7 and 8 (in reality, I only need 1 and 4 to be translated given how I'm deducing the rest)
func determineSimpleNumbers(dict [10]string, pattern []string) [10]string {
	for _, digit := range pattern {
		parsedDigit := string(digit)
		switch len(parsedDigit) {
		case 2:
			dict[1] = parsedDigit
		case 3:
			dict[7] = parsedDigit
		case 4:
			dict[4] = parsedDigit
		case 7:
			dict[8] = parsedDigit
		}
	}
	return dict
}

// how many letters/segments appear in both strings
func countOverlap(a, b string) int {
	count := 0
	for _, char := range b {
		count += strings.Count(a, string(char))
	}
	return count
}

// super dumb approach lmfao
// i'm determining what the output (letter/segment sequence) would translate to, based on what segments
// they overlap with from what we already know (the simple digits 1, 4, 7 and 8)
// length 5: the output for 2, 3 and 5 overlap with output for 4 and 1 uniquely
// length 6: the output for 0, 6 and 9 overlap with output for 4 and 1 uniquely
func determineRemainingNumbers(dict [10]string, pattern []string) [10]string {
	for _, digit := range pattern {
		switch len(digit) {
		case 5:
			if countOverlap(dict[4], digit) == 2 {
				dict[2] = digit
			} else if countOverlap(dict[1], digit) == 1 {
				dict[5] = digit
			} else {
				dict[3] = digit
			}
		case 6:
			if countOverlap(dict[4], digit) == 4 {
				dict[9] = digit
			} else if countOverlap(dict[1], digit) == 1 {
				dict[6] = digit
			} else {
				dict[0] = digit
			}
		}
	}

	return dict
}

// since letters/segments can be shuffled, if they have the same length and equal letters,
// they are the same...
func indexOf(arr [10]string, value string) int {
	for i, v := range arr {
		overlap := countOverlap(v, value)
		if overlap == len(v) && overlap == len(value) {
			return i
		}
	}
	return -1
}

// given a pattern, this will return the translated value. e.g. abdc bcde => 42 (arbitrary example, just to visualize)
func parseOutputValue(dict [10]string, digits []string) int {
	sum := 0
	length := len(digits)
	for index, pwr := 0, float64(length-1); index < length; index, pwr = index+1, pwr-1 {
		digit := digits[index]
		dictValue := indexOf(dict, digit)
		if dictValue == -1 {
			panic("Didn't find translation")
		}
		sum += dictValue * int(math.Pow(10, pwr)) // (value x 10^pwr), e.g. 4 x 10^3 = 4000
	}
	return sum
}

func determineOutputSum(inputFileName string) int {
	inputFile := util.OpenInputFile(8, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	outputSum := 0
	for input.Scan() {
		translationList := [10]string{}
		pattern, outputValue := parseInputLine(input.Text())

		translationList = determineSimpleNumbers(translationList, pattern) // 1, 4, 7 and 8 (since they have unique lengths)
		translationList = determineRemainingNumbers(translationList, pattern)
		outputSum += parseOutputValue(translationList, outputValue)
	}

	return outputSum
}
