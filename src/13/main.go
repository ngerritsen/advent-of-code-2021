package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type fold struct {
	axis string
	n    int
}

//go:embed input.txt
var input string

func main() {
	parts := strings.Split(input, "\n\n")
	coords := parseCoords(parts[0])
	folds := parseFolds(parts[1])

	println(len(doFold(coords, folds[0])))
	printSheet(applyFolds(coords, folds))
}

func applyFolds(coords []coord, folds []fold) []coord {
	for _, f := range folds {
		coords = doFold(coords, f)
	}

	return coords
}

func doFold(coords []coord, f fold) []coord {
	var next []coord

	for _, c := range coords {
		if f.axis == "y" && c.y > f.n {
			c.y = c.y - ((c.y - f.n) * 2)
		} else if f.axis == "x" && c.x > f.n {
			c.x = c.x - ((c.x - f.n) * 2)
		}

		if !hasCoord(next, c) {
			next = append(next, c)
		}
	}

	return next
}

func printSheet(coords []coord) {
	s := coord{x: math.MaxInt, y: math.MaxInt}
	e := coord{x: math.MinInt, y: math.MinInt}

	for _, c := range coords {
		s.x = min(c.x, s.x)
		e.x = max(c.x, e.x)
		s.y = min(c.y, s.y)
		e.y = max(c.y, e.y)
	}

	for y := s.y; y <= e.y; y++ {
		for x := s.x; x <= e.x; x++ {
			symbol := " "

			if hasCoord(coords, coord{x: x, y: y}) {
				symbol = "#"
			}

			print(symbol)
		}
		println()
	}
}

func hasCoord(coords []coord, coord coord) bool {
	for _, c := range coords {
		if c.equals(coord) {
			return true
		}
	}

	return false
}

func parseCoords(input string) []coord {
	var coords []coord
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		pts := strings.Split(l, ",")
		coords = append(coords, coord{
			x: toInt(pts[0]),
			y: toInt(pts[1]),
		})
	}

	return coords
}

func parseFolds(input string) []fold {
	var folds []fold
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		ws := strings.Fields(l)
		pts := strings.Split(ws[2], "=")
		folds = append(folds, fold{
			axis: pts[0],
			n:    toInt(pts[1]),
		})
	}

	return folds
}

func (c *coord) equals(o coord) bool {
	return c.x == o.x && c.y == o.y
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func toInt(str string) int {
	v, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return v
}
