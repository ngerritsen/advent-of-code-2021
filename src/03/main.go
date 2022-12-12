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

	println(getPowerConsumption(lns))
	println(getLifeSupport(lns))
}

func getPowerConsumption(lns []string) int {
	gamma, epsilon := "", ""

	for i := range lns[0] {
		if moreOnes(lns, i) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return binToInt(gamma) * binToInt(epsilon)
}

func getLifeSupport(lns []string) int {
	oxy := narrowDown(lns, true)
	co2 := narrowDown(lns, false)

	return oxy * co2
}

func narrowDown(lns []string, common bool) int {
	for i, _ := range lns[0] {
		var narrowed []string
		oc := moreOnes(lns, i)
		char := "0"[0]

		if common && oc || !common && !oc {
			char = "1"[0]
		}

		for _, ln := range lns {
			if ln[i] == char {
				narrowed = append(narrowed, ln)
			}
		}

		lns = narrowed

		if len(lns) == 1 {
			break
		}
	}

	return binToInt(lns[0])
}

func moreOnes(lns []string, i int) bool {
	var counts [2]int

	for _, ln := range lns {
		counts[int(ln[i])-48]++
	}

	return counts[0] <= counts[1]
}

func binToInt(bin string) int {
	i, _ := strconv.ParseInt(bin, 2, 64)
	return int(i)
}
