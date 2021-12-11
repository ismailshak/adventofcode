// Advent of Code Day 10: Syntax Scoring
package day10

import (
	"aoc/util"
	"bufio"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

var TOKEN_PAIRS_CLOSING = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var TOKEN_PAIRS_OPENING = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var TOKEN_SYNTAX_SCORE = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var TOKEN_AUTOCOMP_SCORE = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func Solve(inputFileName string) {
	fmt.Println("\nSolving Syntax Scoring")
	fmt.Println("----------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	syntaxScore, incompleteStacks := determineSyntaxScore(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", syntaxScore, p1Duration)

	// Part 2
	p2Start := time.Now()
	autoCompScore := determineAutoCompleteScore(incompleteStacks)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", autoCompScore, p2Duration)
}

//------------------------------
// quick Stack implementation, didn't find anything useful in std lib
type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	if len(s) == 0 {
		panic("Can't pop an empty stack")
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Peek() string {
	if len(s) == 0 {
		return ""
	}

	l := len(s)
	return s[l-1]
}

const (
	CYAN  = "\033[36m"
	RESET = "\033[0m"
)

// hard to read the default print because it wraps the output in `[ ]`
func (s Stack) Print() {
	fmt.Print("| ")
	for _, token := range s {
		fmt.Print(CYAN, token, RESET, ", ")
	}
	fmt.Print(" |\n")
}

//-----------------------------

func isOpeningToken(token string) bool {
	_, ok := TOKEN_PAIRS_OPENING[token]
	return ok
}

// returns a bool if line error'd (error is truthy here), the token that caused
// the syntax error, and the stack at the end of "compiling"
func compileLine(line string) (string, bool, Stack) {
	tokens := strings.Split(line, "")
	stack := Stack{}

	for _, token := range tokens {
		// if it's opening, add to the stack since it expects a closing
		if isOpeningToken(token) {
			stack = stack.Push(token)
			continue
		}

		// if we're on a closing token, and the top of the stack is it's opening counterpart,
		// pop the token out of the stack and proceed (i.e. successfully matched a full pair)
		openingPair := TOKEN_PAIRS_CLOSING[token]
		if stack.Peek() == openingPair {
			stack, _ = stack.Pop()
			continue
		}

		// if we get to here, syntax error since the current closing token does not match the
		// opening token the "complier" expected
		return token, true, stack
	}

	// this is an incomplete line (since "all of them" had a syntax error - otherwise these could include
	// syntactically valid lines)
	return "", false, stack
}

func syntaxScore(tokens map[string]int) int {
	sum := 0
	for token, frequency := range tokens {
		sum += TOKEN_SYNTAX_SCORE[token] * frequency
	}
	return sum
}

func determineSyntaxScore(inputFileName string) (int, []Stack) {
	inputFile := util.OpenInputFile(10, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	incompleteStacks := []Stack{}
	syntaxErrorTokens := map[string]int{}

	for input.Scan() {
		errorToken, didError, stack := compileLine(input.Text())
		if didError {
			syntaxErrorTokens[errorToken] += 1
		} else {
			incompleteStacks = append(incompleteStacks, stack)
		}
	}

	return syntaxScore(syntaxErrorTokens), incompleteStacks
}

// creates a completion string from an incomplete stack
func autoComplete(stack Stack) string {
	completionString := ""
	numOfTokens := len(stack)

	for i := 0; i < numOfTokens; i++ {
		var removed string
		stack, removed = stack.Pop()
		completionString += TOKEN_PAIRS_OPENING[removed]
	}

	return completionString
}

func autoCompleteScore(completion string) int {
	sum := 0
	for _, char := range completion {
		sum *= 5
		sum += TOKEN_AUTOCOMP_SCORE[string(char)]
	}
	return sum
}

func determineAutoCompleteScore(incompleteStacks []Stack) int {
	autoCompleteScores := []int{}

	for _, stack := range incompleteStacks {
		completionString := autoComplete(stack)
		autoCompleteScores = append(autoCompleteScores, autoCompleteScore(completionString))
	}

	sort.Ints(autoCompleteScores)
	middleIndex := math.Floor(float64(len(autoCompleteScores) / 2)) // always odd number of lines
	return autoCompleteScores[int(middleIndex)]
}
