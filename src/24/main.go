package main

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const mod = 26

type digits []int
type formula []int

//go:embed example.txt
var code string

func main() {
	f := parseFormula(code)

	println(getMax(f))
	println(getMin(f))
}

func getMax(f formula) string {
	d := makeDigits(len(f)/2, 9)

	for i := 0; i < pow(9, len(d))-1; i++ {
		r := solve(f, d)

		if len(r) > 0 {
			return r
		}

		d.decr()
	}

	return ""
}

func getMin(f formula) string {
	d := makeDigits(len(f)/2, 1)

	for i := 0; i < pow(9, len(d))-1; i++ {
		r := solve(f, d)

		if len(r) > 0 {
			return r
		}

		d.incr()
	}

	return ""
}

func solve(f formula, d digits) string {
	r, z, di := "", 0, 0

	for _, op := range f {
		if op > 0 {
			z = (z * mod) + d[di] + op
			r += strconv.Itoa(d[di])
			di++
		} else {
			i := (z % mod) + op

			if i < 1 || i > 9 {
				return ""
			}

			z = z / mod
			r += strconv.Itoa(i)
		}
	}

	return r
}

func (d digits) decr() {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] > 1 {
			d[i]--
			break
		}

		d[i] = 9
	}
}

func (d digits) incr() {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] < 9 {
			d[i]++
			break
		}

		d[i] = 1
	}
}

func makeDigits(n, v int) digits {
	d := make(digits, n)

	for i := 0; i < n; i++ {
		d[i] = v
	}

	return d
}

func pow(i, e int) int {
	return int(math.Pow(float64(i), float64(e)))
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func parseFormula(input string) formula {
	ops := strings.Split(strings.TrimLeft(input, "inp"), "inp")
	f := make([]int, len(ops))
	rn, rp := regexp.MustCompile(`-\d+`), regexp.MustCompile(`w\nadd y (\d+)`)

	for i, op := range ops {
		if rn.MatchString(op) {
			f[i] = toInt(rn.FindString(op))
		} else {
			f[i] = toInt(rp.FindStringSubmatch(op)[1])
		}
	}

	return f
}
