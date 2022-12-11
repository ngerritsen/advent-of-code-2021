package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lns := strings.Split(input, "\n")

	println(run(lns))
	println(run2(lns))
}

func run(lns []string) int {
	x, y := 0, 0

	for _, ln := range lns {
		cmd, n := parseLine(ln)

		switch cmd {
		case "up":
			y -= n
			break
		case "down":
			y += n
			break
		case "forward":
			x += n
			break
		}
	}

	return x * y
}

func run2(lns []string) int {
	x, y, a := 0, 0, 0

	for _, ln := range lns {
		cmd, n := parseLine(ln)

		switch cmd {
		case "up":
			a -= n
			break
		case "down":
			a += n
			break
		case "forward":
			x += n
			y += n * a
			break
		}
	}

	return x * y
}

func parseLine(ln string) (string, int) {
	pts := strings.Split(ln, " ")
	rn, _ := strconv.ParseInt(pts[1], 10, 64)
	return pts[0], int(rn)
}
