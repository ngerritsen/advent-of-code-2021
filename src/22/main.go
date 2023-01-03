package main

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	x = 0
	y = 1
	z = 2
)

type point [3]int

type command struct {
	on   bool
	cube cube
}

type cube struct {
	from, to point
}

type line struct {
	from, to int
}

type cubes map[cube]bool

func main() {
	cmds := parseCommands(input)

	println(getTotalVol(cmds, cubeFromBound(50)))
	println(getTotalVol(cmds, cubeFromBound(math.MaxInt)))
}

func getTotalVol(cmds []command, bound cube) int {
	cs, v := run(cmds, bound), 0

	for c := range cs {
		v += c.vol()
	}

	return v
}

func run(cmds []command, bound cube) cubes {
	cs := make(cubes)

	for _, cmd := range cmds {
		nc := cmd.cube

		if !nc.intersects(bound) {
			continue
		}

		tmp := make(cubes)

		for c := range cs {
			if !c.intersects(nc) {
				tmp.add(c)
				continue
			}

			for p := range c.split(nc) {
				if !p.intersects(nc) {
					tmp.add(p)
				}
			}
		}

		cs = tmp

		if cmd.on {
			cs.add(nc)
		}
	}

	return cs
}

func cubeFromBound(b int) cube {
	return cube{point{neg(b), neg(b), neg(b)}, point{b, b, b}}
}

func (c cube) intersects(o cube) bool {
	for a := range []int{x, y, z} {
		if !c.edge(a).intersects(o.edge(a)) {
			return false
		}
	}

	return true
}

func (c cube) split(o cube) cubes {
	cs := make(cubes)

	for cx := range c.splitOn(x, o) {
		for cy := range cx.splitOn(y, o) {
			for cz := range cy.splitOn(z, o) {
				cs.add(cz)
			}
		}
	}

	return cs
}

func (c cube) splitOn(a int, o cube) cubes {
	cs := make(cubes)

	if c.from[a] < o.from[a] {
		cs.add(cube{c.from, c.to.with(a, o.from[a]-1)})
	}

	cs.add(cube{
		c.from.with(a, max(c.from[a], o.from[a])),
		c.to.with(a, min(c.to[a], o.to[a])),
	})

	if c.to[a] > o.to[a] {
		cs.add(cube{c.from.with(a, o.to[a]+1), c.to})
	}

	return cs
}

func (c cube) edge(a int) line { return line{c.from[a], c.to[a]} }
func (c cube) vol() int        { return c.edge(x).len() * c.edge(y).len() * c.edge(z).len() }

func (p point) with(a, v int) point {
	p[a] = v
	return p
}

func (l line) len() int               { return l.to - l.from + 1 }
func (l line) intersects(o line) bool { return !l.precedes(o) && !o.precedes(l) }
func (l line) precedes(o line) bool   { return l.from < o.from && l.to < o.from }

func (cs cubes) add(c cube)    { cs[c] = true }
func (cs cubes) remove(c cube) { delete(cs, c) }

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func neg(n int) int {
	return n * -1
}

func min(a, b int) int { return int(math.Min(float64(a), float64(b))) }
func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }

func parseCommands(input string) []command {
	lns := strings.Split(input, "\n")
	cmds := make([]command, len(lns))

	for i, l := range lns {
		pts := regexp.MustCompile(`(-?\d+)`).FindAllString(l, 6)
		cmds[i] = command{
			l[0:2] == "on",
			cube{
				point{toInt(pts[0]), toInt(pts[2]), toInt(pts[4])},
				point{toInt(pts[1]), toInt(pts[3]), toInt(pts[5])},
			},
		}
	}

	return cmds
}
