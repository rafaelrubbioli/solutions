package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func flash(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[x][y] == 0 {
		return 0
	}

	grid[x][y]++
	if grid[x][y] > 9 {
		grid[x][y] = 0
		return 1 + flash(grid, x, y+1) +
			flash(grid, x, y-1) +
			flash(grid, x+1, y) +
			flash(grid, x+1, y+1) +
			flash(grid, x+1, y-1) +
			flash(grid, x-1, y) +
			flash(grid, x-1, y+1) +
			flash(grid, x-1, y-1)
	}

	return 0
}

func countFlashes(grid [][]int, duration int) int {
	flashes := 0
	for step := 0; step < duration; step++ {
		for x := range grid {
			for y := range grid[x] {
				grid[x][y]++
			}
		}

		for x := range grid {
			for y := range grid[x] {
				if grid[x][y] == 10 {
					flashes += flash(grid, x, y)
				}
			}
		}

		fmt.Println("step: ", step)
		for _, line := range grid {
			fmt.Println(line)
		}
	}

	return flashes
}

func calculateSync(grid [][]int, maxDuration int) int {
	for step := 0; step < maxDuration; step++ {
		for x := range grid {
			for y := range grid[x] {
				grid[x][y]++
			}
		}

		stepFlashes := 0
		for x := range grid {
			for y := range grid[x] {
				if grid[x][y] == 10 {
					stepFlashes += flash(grid, x, y)
				}
			}
		}

		fmt.Println("step: ", step)
		for _, line := range grid {
			fmt.Println(line)
		}

		if stepFlashes == len(grid)*len(grid[0]) {
			for _, line := range grid {
				fmt.Println(line)
			}

			return step + 1
		}
	}

	return 0
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	var input [][]int
	for scanner.Scan() {
		line := make([]int, 0)
		for _, c := range scanner.Text() {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}

			line = append(line, n)
		}

		input = append(input, line)
	}

	//fmt.Println(countFlashes(input, 100))
	fmt.Println(calculateSync(input, 1000))
}
