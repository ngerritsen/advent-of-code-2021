package main

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

type target struct{ min, max coord }
type coord struct{ x, y int }

func main() {
	t := parseTarget(input)
	my, hits := getMaxY(t)

	println(my)
	println(hits)
}

func getMaxY(t target) (int, int) {
	my := math.MinInt
	hits := 0

	for vx := getMinVx(t); vx <= t.max.x; vx++ {
		for vy := t.min.y; vy <= t.max.x; vy++ {
			y := simulate(t, vx, vy)

			if y > -1 {
				hits++
			}

			if y > my {
				my = y
			}
		}
	}

	return my, hits
}

func getMinVx(t target) int {
	vx, dx := 0, 0

	for dx < t.min.x {
		vx++
		dx = vx / 2 * (2 + (vx - 1))
	}

	return vx
}

func simulate(t target, vx, vy int) int {
	my, c := 0, coord{0, 0}

	for !pastTarget(t, c) {
		c.x += vx
		c.y += vy

		if c.y > my {
			my = c.y
		}

		if hitTarget(t, c) {
			return my
		}

		if vx > 0 {
			vx--
		}

		vy--
	}

	return math.MinInt
}

func pastTarget(t target, c coord) bool {
	return c.x > t.max.x || c.y < t.min.y
}

func hitTarget(t target, c coord) bool {
	return c.x <= t.max.x && c.x >= t.min.x && c.y <= t.max.y && c.y >= t.min.y
}

func parseTarget(input string) target {
	re := regexp.MustCompile(`(-?\d+)(?:\.{2})(-?\d+)`)
	res := re.FindAllStringSubmatch(input, 4)
	x, y := res[0], res[1]

	return target{
		min: coord{x: toInt(x[1]), y: toInt(y[1])},
		max: coord{x: toInt(x[2]), y: toInt(y[2])},
	}
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
