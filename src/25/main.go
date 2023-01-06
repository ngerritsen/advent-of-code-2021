package main

import (
	_ "embed"
	"strings"
)

const (
	east  = '>'
	south = 'v'
	none  = '.'
)

type row []rune
type grid []row

//go:embed input.txt
var input string

func main() {
	g := parseGrid(input)

	println(g.settle())
}

func (g grid) settle() int {
	s, ng, moved := 0, g, true

	for moved {
		ng, moved = ng.next()
		s++
	}

	return s
}

func (g grid) next() (grid, bool) {
	g, me := g.east()
	g, ms := g.south()

	return g, me || ms
}

func (g grid) east() (grid, bool) {
	ng := makeGrid(g.width(), g.height())
	m := false

	for y, r := range g {
		for x := range r {
			v := g[y][x]

			if v == none {
				continue
			}

			nx := (x + 1) % g.width()

			if v == east && g[y][nx] == none {
				ng[y][nx] = v
				m = true
			} else {
				ng[y][x] = v
			}
		}
	}

	return ng, m
}

func (g grid) south() (grid, bool) {
	ng := makeGrid(g.width(), g.height())
	m := false

	for y, r := range g {
		for x := range r {
			v := g[y][x]

			if v == none {
				continue
			}

			ny := (y + 1) % g.height()

			if v == south && g[ny][x] == none {
				ng[ny][x] = v
				m = true
			} else {
				ng[y][x] = v
			}
		}
	}

	return ng, m
}

func (g grid) width() int {
	return len(g[0])
}

func (g grid) height() int {
	return len(g)
}

func makeGrid(w, h int) grid {
	g := make(grid, h)

	for y := range g {
		g[y] = make(row, w)

		for x := range g[y] {
			g[y][x] = none
		}
	}

	return g
}

func parseGrid(input string) grid {
	lns := strings.Split(input, "\n")
	g := makeGrid(len(lns[0]), len(lns))

	for y, l := range lns {
		for x, c := range l {
			g[y][x] = c
		}
	}

	return g
}
