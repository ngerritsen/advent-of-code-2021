package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type grid [][]int

func main() {
	lines := strings.Split(input, "\n")
	grid := makeGrid(lines)
	risk, basins := getStats(grid)

	println(risk)
	println(getBasinScore(basins))
}

func getStats(grid grid) (int, []int) {
	var basins []int
	risk := 0

	for y, row := range grid {
		for x, z := range row {
			if (y == 0 || grid[y-1][x] > z) &&
				(y == len(grid)-1 || grid[y+1][x] > z) &&
				(x == 0 || grid[y][x-1] > z) &&
				(x == len(row)-1 || grid[y][x+1] > z) {
				risk += z + 1
				basins = append(basins, measureBasin(grid, x, y, 0))
			}
		}
	}

	return risk, basins
}

func getBasinScore(basins []int) int {
	sort.Ints(basins)

	tot := 1

	for _, s := range basins[len(basins)-3:] {
		tot *= s
	}

	return tot
}

func measureBasin(grid grid, x, y, s int) int {
	if grid[y][x] == 9 {
		return s
	}

	grid[y][x] = 9

	if y > 0 {
		s = measureBasin(grid, x, y-1, s)
	}

	if y < len(grid)-1 {
		s = measureBasin(grid, x, y+1, s)
	}

	if x > 0 {
		s = measureBasin(grid, x-1, y, s)
	}

	if x < len(grid[0])-1 {
		s = measureBasin(grid, x+1, y, s)
	}

	return s + 1
}

func makeGrid(lines []string) grid {
	grid := make(grid, len(lines))

	for y, l := range lines {
		grid[y] = make([]int, len(l))

		for x, zs := range strings.Split(l, "") {
			z, _ := strconv.Atoi(zs)
			grid[y][x] = z
		}
	}

	return grid
}
