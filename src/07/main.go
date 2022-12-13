package main

import (
	_ "embed"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type costs = map[int]int
type crabs = []int

func main() {
	strs := strings.Split(input, ",")
	crabs := make(crabs, len(strs))

	for i, str := range strs {
		crabs[i] = toInt(str)
	}

	println(findOptimum(crabs, false))
	println(findOptimum(crabs, true))
}

func findOptimum(crabs crabs, useCost bool) int {
	max := getMax(crabs)
	fuelCosts := getFuelCosts(max)
	x := getMedian(crabs)
	cost := getCost(crabs, x, useCost, fuelCosts)

	for true {
		l := getCost(crabs, x-1, useCost, fuelCosts)
		if l < cost {
			x--
			cost = l
			continue
		}

		r := getCost(crabs, x+1, useCost, fuelCosts)
		if r < cost {
			x++
			cost = r
			continue
		}

		break
	}

	return cost
}

func getMedian(crabs crabs) int {
	sorted := make([]int, len(crabs))
	copy(sorted, crabs)
	sort.Ints(sorted)
	return sorted[len(sorted)/2]
}

func getFuelCosts(max int) costs {
	fuelCosts := make(costs)
	prev := 0

	for i := 0; i <= max; i++ {
		fuelCosts[i] = prev + i
		prev = fuelCosts[i]
	}

	return fuelCosts
}

func getMax(crabs crabs) int {
	max := 0

	for _, c := range crabs {
		if c > max {
			max = c
		}
	}

	return max
}

func getCost(crabs crabs, to int, useCost bool, fuelCosts costs) int {
	tot := 0

	for _, c := range crabs {
		cost := int(math.Abs(float64(c - to)))

		if useCost {
			cost = fuelCosts[cost]
		}

		tot += cost
	}

	return tot
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		log.Fatalln(err)
	}

	return n
}
