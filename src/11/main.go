package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type grid [][]int

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	println(countTotalFlashes(parseGrid(lines)))
	println(getFullFlashStep(parseGrid(lines)))
}

func getFullFlashStep(grid grid) int {
	s := len(grid) * len(grid[0])
	i := 0

	for countFlashes(grid) != s {
		i++
		incrGrid(grid)
		processFlashes(grid)
	}

	return i
}

func countTotalFlashes(grid grid) int {
	flashes := 0

	for i := 0; i < 100; i++ {
		incrGrid(grid)
		processFlashes(grid)
		flashes += countFlashes(grid)
	}

	return flashes
}

func processFlashes(grid grid) {
	for y, row := range grid {
		for x, _ := range row {
			if row[x] > 9 {
				flash(grid, x, y)
			}
		}
	}
}

func flash(grid grid, x, y int) {
	w := len(grid)
	h := len(grid[0])
	grid[y][x] = 0

	for yy := max(y-1, 0); yy <= min(y+1, h-1); yy++ {
		for xx := max(x-1, 0); xx <= min(x+1, w-1); xx++ {
			if grid[yy][xx] == 0 {
				continue
			}

			grid[yy][xx]++

			if grid[yy][xx] > 9 {
				flash(grid, xx, yy)
			}
		}
	}
}

func incrGrid(grid grid) {
	for _, row := range grid {
		for x, _ := range row {
			row[x]++
		}
	}
}

func countFlashes(grid grid) int {
	flashes := 0

	for _, row := range grid {
		for x, _ := range row {
			if row[x] == 0 {
				flashes += 1
			}
		}
	}

	return flashes
}

func parseGrid(lines []string) grid {
	grid := make(grid, len(lines))

	for y, line := range lines {
		for _, c := range line {
			grid[y] = append(grid[y], parseInt(string(c)))
		}
	}

	return grid
}

func parseInt(str string) int {
	val, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return val
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
