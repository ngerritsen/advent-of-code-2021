package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type pixels map[coord]bool
type algorithm map[int]bool

type image struct {
	pixels pixels
	alg    algorithm
	size   int
}

//go:embed input.txt
var input string

func main() {
	img := parseInput(input)

	println(img.enhance(2).crop().lit())
	println(img.enhance(50).crop().lit())
}

func getBounds(size int) (int, int) {
	e, s := size/2, size/-2

	return s, e
}

func (img image) lit() int {
	return len(img.pixels)
}

func (img image) crop() image {
	s, e := getBounds(img.size)
	px := make(pixels)

	for p := range img.pixels {
		if p.x >= s && p.x <= e && p.y >= s && p.y <= e {
			px.add(p)
		}
	}

	return image{px, img.alg, img.size}
}

func (img image) enhance(n int) image {
	px := make(pixels)
	s, e := getBounds(img.size + (n * 3))

	for y := s; y <= e; y++ {
		for x := s; x <= e; x++ {
			c := coord{x, y}

			if img.isLit(c) {
				px.add(c)
			}
		}
	}

	nextImg := image{px, img.alg, img.size + 2}

	if n > 1 {
		return nextImg.enhance(n - 1)
	}

	return nextImg
}

func (img image) isLit(c coord) bool {
	s := ""

	for y := c.y - 1; y <= c.y+1; y++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			v := '0'

			if img.pixels.has(coord{x, y}) {
				v = '1'
			}

			s += string(v)
		}
	}

	val, _ := strconv.ParseInt(s, 2, 0)

	return img.alg.has(int(val))
}

func (px pixels) has(c coord) bool { return px[c] }
func (px pixels) add(c coord)      { px[c] = true }

func (alg algorithm) has(n int) bool { return alg[n] }
func (alg algorithm) add(n int)      { alg[n] = true }

func parseInput(input string) image {
	parts := strings.Split(input, "\n\n")
	px, size := parsePixels(parts[1])
	alg := parseAlgorithm(parts[0])

	return image{px, alg, size}
}

func parsePixels(input string) (pixels, int) {
	p := make(pixels)
	lines := strings.Split(input, "\n")
	o := len(lines) / 2

	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				p[coord{x - o, y - o}] = true
			}
		}
	}

	return p, len(lines)
}

func parseAlgorithm(input string) algorithm {
	alg := make(algorithm)

	for n, c := range input {
		if c == '#' {
			alg[n] = true
		}
	}

	return alg
}
