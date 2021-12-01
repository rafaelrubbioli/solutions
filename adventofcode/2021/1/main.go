package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func increaseCount(input []int) int {
	count := 0
	lastSeen := -1
	for _, val := range input {
		if lastSeen == -1 {
			lastSeen = val
			continue
		}

		if val > lastSeen {
			count++
		}

		lastSeen = val
	}

	return count
}

type window struct {
	sum      int
	elements []int
}

func windowIncreaseCount(input []int) int {
	count := 0
	window := window{
		sum:      0,
		elements: make([]int, 0, 3),
	}

	for _, val := range input {
		window.elements = append(window.elements, val)
		window.sum += val

		if len(window.elements) > 3 {
			oldSum := window.sum - window.elements[3]
			window.sum -= window.elements[0]
			if window.sum > oldSum {
				count++
			}

			window.elements = window.elements[1:]
		}
	}

	return count
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanWords)

	input := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, num)
	}

	part1 := increaseCount(input)
	fmt.Println("Part 1:", part1)

	part2 := windowIncreaseCount(input)
	fmt.Println("Part 2:", part2)
}
