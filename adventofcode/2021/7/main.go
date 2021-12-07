package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minAndMax(items []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for _, item := range items {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}

	return min, max
}

func mean(positions []int) int {
	sum := 0
	for _, v := range positions {
		sum += v
	}

	return sum / len(positions)
}

func median(positions []int) int {
	sort.Ints(positions)
	if len(positions)%2 == 0 {
		return (positions[len(positions)/2-1] + positions[len(positions)/2]) / 2
	}

	return positions[len(positions)/2]
}

func fuelToAlignConstant(positions []int, target int) int {
	fuel := 0
	for _, position := range positions {
		fuel += int(math.Abs(float64(position - target)))
	}

	return fuel
}

func fuelToAlignIncreasing(positions []int, target int) int {
	fuel := 0
	for _, position := range positions {
		dist := int(math.Abs(float64(position - target)))
		fuel += dist * (dist + 1) / 2
	}

	return fuel
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	positions := make([]int, 0)
	input := strings.Replace(string(content), "\n", "", -1)
	for _, position := range strings.Split(input, ",") {
		n, err := strconv.Atoi(position)
		if err != nil {
			log.Fatal(err)
		}

		positions = append(positions, n)
	}

	median := median(positions)
	fmt.Println("Part 1: ", fuelToAlignConstant(positions, median))

	min, max := minAndMax(positions)
	lowest := math.MaxInt64
	for i := min; i <= max; i++ {
		current := fuelToAlignIncreasing(positions, i)
		if current < lowest {
			lowest = current
		}
	}

	fmt.Println("Part 2: ", lowest)
}
