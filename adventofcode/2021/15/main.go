package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

func getID(i, j int) int {
	return i*1000 + j
}

func getShortestPath(grid [][]int) (int, error) {
	graph := dijkstra.NewGraph()
	for i, row := range grid {
		for j := range row {
			graph.AddVertex(getID(i, j))
		}
	}

	for i, row := range grid {
		for j, val := range row {
			id := getID(i, j)
			risk := int64(val)
			if i > 0 {
				err := graph.AddArc(getID(i-1, j), id, risk)
				if err != nil {
					return 0, err
				}
			}

			if i < len(grid)-1 {
				err := graph.AddArc(getID(i+1, j), id, risk)
				if err != nil {
					return 0, err
				}
			}

			if j > 0 {
				err := graph.AddArc(getID(i, j-1), id, risk)
				if err != nil {
					return 0, err
				}
			}

			if j < len(row)-1 {
				err := graph.AddArc(getID(i, j+1), id, risk)
				if err != nil {
					return 0, err
				}
			}
		}
	}

	path, err := graph.Shortest(getID(0, 0), getID(len(grid)-1, len(grid[0])-1))
	if err != nil {
		return 0, err
	}

	return int(path.Distance), nil
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	grid := make([][]int, 0)
	for scanner.Scan() {
		line := make([]int, 0)
		for _, c := range scanner.Text() {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}

			line = append(line, n)
		}

		grid = append(grid, line)
	}

	shortestPath, err := getShortestPath(grid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", shortestPath)
	rows, cols := len(grid), len(grid[0])

	newGrid := make([][]int, len(grid)*5)
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid[0])*5)
	}

	for i, row := range grid {
		for j, val := range row {
			for rowMult := 0; rowMult < 5; rowMult++ {
				for colMult := 0; colMult < 5; colMult++ {
					newValue := val + rowMult + colMult
					if newValue > 9 {
						newValue -= 9
					}

					newGrid[i+(rowMult*rows)][j+(colMult*cols)] = newValue
				}
			}
		}
	}

	shortestPath2, err := getShortestPath(newGrid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2: ", shortestPath2)
}
