package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lns := strings.Split(input, "\n")
	depths := toInts(lns)

	println(depthIncr(depths))
	println(depthIncr(getWindows(depths)))
}

func getWindows(depths []int) []int {
	l := len(depths)
	windows := make([]int, l-2)

	for i, d := range depths {
		if i > 1 {
			windows[i-2] += d
		}
		if i > 0 && i < l-1 {
			windows[i-1] += d
		}
		if i < l-2 {
			windows[i] += d
		}
	}

	return windows
}

func depthIncr(depths []int) int {
	tot := 0
	prev := math.MaxInt

	for _, d := range depths {
		if d > prev {
			tot++
		}

		prev = d
	}

	return tot
}

func toInts(strs []string) []int {
	res := make([]int, len(strs))

	for i, l := range strs {
		v, _ := strconv.ParseInt(l, 10, 64)
		res[i] = int(v)
	}

	return res
}
