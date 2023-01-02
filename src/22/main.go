package main

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed example.txt
var input string

type coord struct {
	x, y, z int
}

type command struct {
	on       bool
	from, to coord
}

type cubes map[coord]bool

func main() {
	cmds := parseCommands(input)

	println(len(run(cmds, -50, 50)))
}

func run(cmds []command, s, e int) cubes {
	cbs := make(cubes)

	for i, cmd := range cmds {
		if !isRelevant(cmd, s, e) {
			continue
		}

		for x := max(cmd.from.x, s); x <= min(cmd.to.x, e); x++ {
			for y := max(cmd.from.y, s); y <= min(cmd.to.y, e); y++ {
				for z := max(cmd.from.z, s); z <= min(cmd.to.z, e); z++ {
					if cmd.on {
						cbs.add(coord{x, y, z})
					} else {
						cbs.remove(coord{x, y, z})
					}
				}
			}
		}
	}

	return cbs
}

func parseCommands(input string) []command {
	lns := strings.Split(input, "\n")
	cmds := make([]command, len(lns))

	for i, l := range lns {
		pts := regexp.MustCompile(`(-?\d+)`).FindAllString(l, 6)
		cmds[i] = command{
			on:   l[0:2] == "on",
			from: coord{toInt(pts[0]), toInt(pts[2]), toInt(pts[4])},
			to:   coord{toInt(pts[1]), toInt(pts[3]), toInt(pts[5])},
		}
	}

	return cmds
}

func isRelevant(cmd command, s, e int) bool {
	return hasOverlap(cmd.from.x, cmd.to.x, s, e) &&
		hasOverlap(cmd.from.y, cmd.to.y, s, e) &&
		hasOverlap(cmd.from.z, cmd.to.z, s, e)
}

func hasOverlap(as, ae, bs, be int) bool {
	return isBetween(as, bs, be) || isBetween(ae, bs, be)
}

func isBetween(m, s, e int) bool {
	return s <= m && m <= e
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func (cbs cubes) add(c coord) {
	cbs[c] = true
}

func (cbs cubes) remove(c coord) {
	delete(cbs, c)
}
