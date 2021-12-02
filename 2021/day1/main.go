// AdventOfCode Day 1: Sonar Sweep.
// Choosing to stream the data in, versus swallowing the whole input in-memory.
// Choosing to exercise approaches I don't usually take in a professional setting, lol
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Solving Sonar Sweep")
	start := time.Now()
	fmt.Printf("Part 1 Result: %v\n", determineDepth())
	p1Duration := time.Since(start)
	fmt.Printf("Part 2 Result: %v\n", determineCycledDepth())
	p2Duration := time.Since(start)
	fmt.Println("---------------------")
	fmt.Printf("Part 1 took %v. Part 2 took %v\n", p1Duration, p2Duration)
}

func openInputFile() *os.File {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic("Error. Failed to read input file.")
	}
	return inputFile
}

// Why does .ParseInt always return int64? :thinking:
func parseInt(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 32)
}

func determineDepth() int32 {
	var numberOfIncreases int32

	inputFile := openInputFile()
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	// Scanning once to skip over the first token/input
	input.Scan()
	previous, _ := parseInt(input.Text())

	for input.Scan() {
		current, _ := parseInt(input.Text())
		if current > previous {
			numberOfIncreases++
		}
		previous = current
	}

	return numberOfIncreases
}

type DepthNode struct {
	depth int64
	next  *DepthNode
}

func (node *DepthNode) isGreater(inNode *DepthNode) bool {
	return node.depth > inNode.depth
}

type LinkedList struct {
	length int8
	head   *DepthNode
	tail   *DepthNode
}

// a+b+c < b+c+d == a < d
func (list *LinkedList) isNewDepthGreater() bool {
	return list.tail.isGreater(list.head)
}

func (list *LinkedList) addNode(value int64) {
	depthNode := DepthNode{
		depth: value,
	}

	if list.length == 0 {
		list.head = &depthNode
		list.tail = &depthNode
	} else {
		list.tail.next = &depthNode
		list.tail = &depthNode
	}

	list.length++
}

func (list *LinkedList) removeHead() {
	if list.head == nil || list.length == 0 {
		panic("Error. Tried to remove head from an empty linked list")
	}

	newHead := list.head.next
	list.head = newHead
	list.length--
}

/*
	If we're trying to determine if a+b+c < b+c+d then all we need to check for
	is if a < d (since b and c cancel out there). Using a linked list because it
	fits the whole "i only care about the first and last elements inserted" approach
	I'm going for. If I make a linked list and keep it at length 4, I can cycle
	the "three-measured" depths in and out of the list. So the above structs 'n methods
	are for that (I don't really need to store b and c, but what the heck)
*/
func determineCycledDepth() int32 {
	var numberOfIncreases int32

	inputFile := openInputFile()
	defer inputFile.Close()

	depthList := LinkedList{}
	input := bufio.NewScanner(inputFile)

	for input.Scan() {
		depthValue, _ := parseInt(input.Text())
		depthList.addNode(depthValue)
		if depthList.length == 4 {
			if depthList.isNewDepthGreater() {
				numberOfIncreases++
			}
			depthList.removeHead()
		}
	}

	return numberOfIncreases
}
