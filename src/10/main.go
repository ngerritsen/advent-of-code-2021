package main

import (
	_ "embed"
	"errors"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

var corruptPoints = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var completePoints = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
var pairs = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

func main() {
	lines := strings.Split(input, "\n")

	println(scoreCorrupt(lines))
	println(scoreIncomplete(lines))
}

func scoreIncomplete(lines []string) int {
	var scores []int

	for _, line := range lines {
		_, stack, err := tryParse(line)

		if err != nil {
			continue
		}

		var score int

		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			score += completePoints[pairs[stack[i]]]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func scoreCorrupt(lines []string) int {
	scores := map[rune]int{')': 0, ']': 0, '}': 0, '>': 0}
	total := 0

	for _, line := range lines {
		c, _, err := tryParse(line)

		if err != nil {
			scores[c] += 1
		}
	}

	for k, v := range scores {
		total += corruptPoints[k] * v
	}

	return total
}

func tryParse(line string) (rune, []rune, error) {
	var stack []rune
	var bug rune

	for _, c := range []rune(line) {
		_, opens := pairs[c]

		if opens {
			stack = append(stack, c)
		} else if pairs[stack[len(stack)-1]] != c {
			return c, stack, errors.New("corrupt")
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return bug, stack, nil
}
