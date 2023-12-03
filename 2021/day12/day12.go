package day12

import (
	"aoc/util"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Passage Pathing")
	fmt.Println("-----------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineTotalPaths(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	//// Part 2
	//p2Start := time.Now()
	//p2Result := determineX()
	//p2Duration := time.Since(p2Start)
	//fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

type Vertex struct {
	adjacent []string
}

func (v *Vertex) AddAdjacent(name string) *Vertex {
	v.adjacent = append(v.adjacent, name)
	return v
}

type Graph struct {
	vertices map[string]*Vertex
}

func (g *Graph) Add(to, from string) {
	if !g.Exists(from) {
		g.vertices[from] = &Vertex{}
	}
	g.vertices[from] = g.vertices[from].AddAdjacent(to)
}

func (g *Graph) Exists(name string) bool {
	_, ok := g.vertices[name]
	return ok
}

func parseLine(line string) (string, string) {
	parts := strings.Split(line, "-")
	return parts[0], parts[1]
}

func isUpperCase(v string) bool {
	return v == strings.ToUpper(v)
}

func indexOf(list []string, v string) int {
	for i, e := range list {
		if e == v {
			return i
		}
	}
	return -1
}

func countPaths(graph *Graph) int {
	pathCount := 0
	stack := util.Stack{[]string{"start"}} // initial state: [["start"]]

	for !stack.IsEmpty() {
		var path []string
		stack, path = stack.Pop()
		currentAdjacency := graph.vertices[path[len(path)-1]].adjacent

		//fmt.Println("path:", path, "- node:", path[len(path)-1], "- adj list:", currentAdjacency)
		for _, adj := range currentAdjacency {
			fmt.Println(adj)
			if adj == "end" {
				fmt.Println("reached end", stack, "- node:", adj, "- path:", path)
				pathCount++
				continue
			}

			if !isUpperCase(adj) {
				if indexOf(path, adj) == -1 {
					stack = stack.Push(append(path, adj))
					fmt.Println("small cave (nv):", adj, "- stack:", stack)
					continue
				}
				//fmt.Println("small cave (v):", adj, "- stack:", stack)
				continue
			}

			stack = stack.Push(append(path, adj))
			fmt.Println("big cave (v):", adj, "- stack:", stack)
		}
	}

	return pathCount
}

func determineTotalPaths(inputFileName string) int {
	inputFile := util.OpenInputFile(12, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	graph := Graph{vertices: make(map[string]*Vertex)}

	for input.Scan() {
		from, to := parseLine(input.Text())
		graph.Add(to, from)
		graph.Add(from, to)
	}

	return countPaths(&graph)
}
