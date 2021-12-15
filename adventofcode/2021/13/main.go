package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func foldUp(grid [][]int, y int) [][]int {
	resultGrid := newGrid(y, len(grid[0]))
	newSize := len(grid) - 1

	for i, row := range grid[:y] {
		for j := range row {
			resultGrid[i][j] = grid[i][j] | grid[newSize-i][j]
		}
	}

	return resultGrid
}

func foldLeft(grid [][]int, x int) [][]int {
	resultGrid := newGrid(len(grid), x)
	newSize := len(grid[0]) - 1

	for i := range grid {
		for j := range grid[i][:x] {
			resultGrid[i][j] = grid[i][j] | grid[i][newSize-j]
		}
	}

	return resultGrid
}

func followInstructions(grid [][]int) [][]int {
	grid = foldLeft(grid, 655)
	grid = foldUp(grid, 447)
	grid = foldLeft(grid, 327)
	grid = foldUp(grid, 223)
	grid = foldLeft(grid, 163)
	grid = foldUp(grid, 111)
	grid = foldLeft(grid, 81)
	grid = foldUp(grid, 55)
	grid = foldLeft(grid, 40)
	grid = foldUp(grid, 27)
	grid = foldUp(grid, 13)
	grid = foldUp(grid, 6)
	return grid
}

func countDots(grid [][]int) int {
	grid = foldLeft(grid, 655)
	count := 0
	for _, row := range grid {
		for _, item := range row {
			if item == 1 {
				count++
			}
		}
	}

	return count
}

func print(grid [][]int) {
	fmt.Println("=========================")
	for _, row := range grid {
		for _, item := range row {
			if item == 1 {
				fmt.Print("#")
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println("=========================")
}

func newGrid(xSize, ySize int) [][]int {
	grid := make([][]int, xSize)
	for i := 0; i < xSize; i++ {
		grid[i] = make([]int, ySize)
	}

	return grid
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	maxX, maxY := 0, 0
	points := make([][2]int, 0)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		if x > maxX {
			maxX = x
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		if y > maxY {
			maxY = y
		}

		points = append(points, [2]int{x, y})
	}

	grid := newGrid(maxY+1, maxX+1)

	for _, point := range points {
		grid[point[1]][point[0]] = 1
	}

	//fmt.Println("Part 1: ", countDots(grid))
	resultGrid := followInstructions(grid)
	fmt.Println("Part 2: ")
	print(resultGrid)
}
