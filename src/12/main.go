package main

import (
	_ "embed"
	"strings"
)

type cave map[string][]string
type stack []state
type state struct {
	node    string
	visited []string
	twice   bool
}

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	cave := parseCave(lines)

	println(getPaths(cave, false))
	println(getPaths(cave, true))
}

func getPaths(cave cave, twice bool) int {
	var stack stack
	paths := 0
	stack = stack.push(state{
		node:  "start",
		twice: !twice,
	})

	for len(stack) > 0 {
		var s state
		stack, s = stack.pop()

		if isVisited(s.visited, s.node) {
			if s.twice || s.node == "start" {
				continue
			}

			s.twice = true
		}

		if s.node == "end" {
			paths++
			continue
		}

		if isLower(s.node) {
			s.visited = append(s.visited, s.node)
		}

		for _, node := range cave[s.node] {
			stack = append(stack, state{
				node:    node,
				visited: s.visited,
				twice:   s.twice,
			})
		}
	}

	return paths
}

func isVisited(visited []string, node string) bool {
	for _, n := range visited {
		if n == node {
			return true
		}
	}

	return false
}

func isLower(str string) bool {
	return strings.ToLower(str) == str
}

func (stack *stack) push(st state) stack {
	return append(*stack, st)
}

func (stack *stack) pop() (stack, state) {
	st := *stack
	l := len(st) - 1
	s := st[l]
	return st[:l], s
}

func parseCave(lines []string) cave {
	cave := make(cave)

	for _, line := range lines {
		parts := strings.Split(line, "-")

		for i, node := range parts {
			neighbour := parts[(i+1)%2]

			_, exists := cave[node]

			if !exists {
				cave[node] = make([]string, 0)
			}

			cave[node] = append(cave[node], neighbour)
		}
	}

	return cave
}
