package main

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

type grid [][]int
type coord struct {
	x int
	y int
}

//go:embed input.txt
var input string

func main() {
	lns := strings.Split(input, "\n")
	crds := parseCoords(lns)
	grid := makeGrid(getMax(crds))
	grid = drawStraightVents(grid, crds)

	println(getTotOverlap(grid))

	grid = drawDiagonalVents(grid, crds)

	println(getTotOverlap(grid))
}

func getTotOverlap(grid grid) int {
	var tot int

	for _, row := range grid {
		for _, col := range row {
			if col > 1 {
				tot++
			}
		}
	}

	return tot
}

func drawStraightVents(grid grid, crds [][]coord) grid {
	for _, l := range crds {
		a, b := l[0], l[1]

		if a.x == b.x {
			s, e := order(a.y, b.y)

			for y := s; y <= e; y++ {
				grid[y][a.x]++
			}
		} else if a.y == b.y {
			s, e := order(a.x, b.x)

			for x := s; x <= e; x++ {
				grid[a.y][x]++
			}
		}
	}

	return grid
}

func drawDiagonalVents(grid grid, crds [][]coord) grid {
	for _, l := range crds {
		a, b := l[0], l[1]

		if a.x == b.x || a.y == b.y {
			continue
		}

		s, e := a, b
		asc := true

		if a.x > b.x {
			s, e = b, a
		}

		if s.y > e.y {
			asc = false
		}

		for x := s.x; x <= e.x; x++ {
			d := x - s.x

			if !asc {
				d *= -1
			}

			grid[s.y+d][x]++
		}
	}

	return grid
}

func getMax(coords [][]coord) (int, int) {
	x, y := 0, 0

	for _, crds := range coords {
		for _, c := range crds {
			x = max(x, c.x)
			y = max(y, c.y)
		}
	}

	return x, y
}

func parseCoords(lns []string) [][]coord {
	var crds [][]coord

	for _, ln := range lns {
		coords := strings.Split(ln, " -> ")
		a, b := parseCoord(coords[0]), parseCoord(coords[1])
		crds = append(crds, []coord{a, b})
	}

	return crds
}

func makeGrid(wi int, hi int) grid {
	var grid = make([][]int, hi+1)

	for i := range grid {
		grid[i] = make([]int, wi+1)
	}

	return grid
}

func order(a int, b int) (int, int) {
	if a > b {
		return b, a
	}

	return a, b
}

func parseCoord(str string) coord {
	pts := strings.Split(str, ",")
	return coord{x: toInt(pts[0]), y: toInt(pts[1])}
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		log.Fatalln(err)
	}

	return n
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
