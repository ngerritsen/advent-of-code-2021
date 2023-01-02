package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var orientationCache = make(map[coord][]coord)

const overlap = 12

type coord struct {
	x, y, z int
}

type scanner struct {
	id      int
	pos     coord
	beacons []coord
}

type result struct {
	offset      coord
	orientation int
}

type pair struct {
	a, b int
}

type ScannerMap map[int]*scanner

func main() {
	scanners := parseScanners(input)
	scanners = normalize(scanners)

	println(getTotalUniqueBeacons(scanners))
	println(getMaxScannerDist(scanners))
}

func getTotalUniqueBeacons(scanners ScannerMap) int {
	all := make(map[coord]int)

	for _, s := range scanners {
		for _, b := range s.beacons {
			all[b]++
		}
	}

	return len(all)
}

func getMaxScannerDist(scanners ScannerMap) int {
	md := math.MinInt

	for i := 0; i < len(scanners)-2; i++ {
		for j := i + 1; j < len(scanners)-1; j++ {
			d := scanners[i].pos.dist(scanners[j].pos)

			if d > md {
				md = d
			}
		}
	}

	return md
}

func normalize(scanners ScannerMap) ScannerMap {
	located := make(ScannerMap)
	located[0] = scanners[0]
	done := make(map[pair]bool)
	delete(scanners, 0)

	for len(scanners) > 0 {
		for _, l := range located {
			for _, s := range scanners {
				p := pair{l.id, s.id}

				if done[p] {
					continue
				}

				m := match(l.beacons, s.beacons, overlap)

				done[p] = true

				if m != nil {
					normalized := transform(s.beacons, m.orientation, m.offset)
					located[s.id] = &scanner{id: s.id, pos: m.offset, beacons: normalized}
					delete(scanners, s.id)
				}
			}
		}
	}

	return located
}

func transform(coords []coord, o int, offset coord) []coord {
	transformed := make([]coord, len(coords))

	for i, c := range coords {
		transformed[i] = getOrientations(c)[o].add(offset)
	}

	return transformed
}

func match(a, b []coord, n int) *result {
	for o := range getOrientations(coord{0, 0, 0}) {
		offsets := make(map[coord]int)

		for _, bb := range b {
			bbo := getOrientations(bb)[o]

			for _, ba := range a {
				d := bbo.diff(ba)
				offsets[d]++

				if offsets[d] == n {
					return &result{d, o}
				}
			}
		}
	}

	return nil
}

func getOrientations(c coord) []coord {
	cached, hit := orientationCache[c]

	if hit {
		return cached
	}

	var orientations []coord
	x, y, z := c.x, c.y, c.z

	directions := []coord{
		{x, y, z},
		{neg(z), y, x},
		{neg(x), y, neg(z)},
		{z, y, neg(x)},
		{x, neg(z), y},
		{x, z, neg(y)},
	}

	for _, d := range directions {
		orientations = append(orientations, getRotations(d)...)
	}

	orientationCache[c] = orientations

	return orientations
}

func getRotations(c coord) []coord {
	x, y, z := c.x, c.y, c.z

	return []coord{
		{x, y, z},
		{neg(y), x, z},
		{neg(x), neg(y), z},
		{y, neg(x), z},
	}
}

func parseScanners(data string) ScannerMap {
	chunks := strings.Split(data, "\n\n")
	scanners := make(ScannerMap)

	for id, chunk := range chunks {
		s := new(scanner)
		s.id = id

		for _, l := range strings.Split(chunk, "\n")[1:] {
			p := strings.Split(l, ",")
			s.beacons = append(s.beacons, coord{toInt(p[0]), toInt(p[1]), toInt(p[2])})
		}

		scanners[id] = s
	}

	return scanners
}

func (a coord) diff(b coord) coord {
	return coord{b.x - a.x, b.y - a.y, b.z - a.z}
}

func (a coord) dist(b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

func (a coord) add(b coord) coord {
	return coord{b.x + a.x, b.y + a.y, b.z + a.z}
}

func (a coord) equals(b coord) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

func neg(n int) int {
	return n * -1
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func toInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
