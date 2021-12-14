package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//	start-A
//	start-b
//	A-c
//	A-b
//	b-d
//	A-end
//	b-end

//		start
//		/   \
//	c--A-----b--d
//		\   /
//		end

func isSmall(s string) bool {
	return strings.ToUpper(s) != s
}

func copy(input map[string]bool) map[string]bool {
	output := make(map[string]bool)
	for k, v := range input {
		output[k] = v
	}

	return output
}

func traverse(paths map[string][]string, node string, visited map[string]bool) [][]string {
	results := make([][]string, 0)
	for _, next := range paths[node] {
		if next == "end" {
			results = append(results, []string{next})
			continue
		}

		visitedForNext := copy(visited)
		if isSmall(next) {
			if visited[next] {
				continue
			}

			visitedForNext[next] = true
		}

		for _, result := range traverse(paths, next, visitedForNext) {
			results = append(results, append([]string{next}, result...))
		}
	}

	return results
}

func traverseWithDoubleVisit(paths map[string][]string, node string, visited map[string]bool, hasDoubleVisit bool) [][]string {
	results := make([][]string, 0)
	for _, next := range paths[node] {
		if next == "end" {
			results = append(results, []string{next})
			continue
		}

		visitedForNext := copy(visited)
		doubleVisit := hasDoubleVisit
		if isSmall(next) {
			if visited[next] {
				if hasDoubleVisit {
					continue
				}

				doubleVisit = true
			}

			visitedForNext[next] = true
		}

		for _, result := range traverseWithDoubleVisit(paths, next, visitedForNext, doubleVisit) {
			results = append(results, append([]string{next}, result...))
		}
	}

	return results
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	paths := make(map[string][]string)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		switch {
		case parts[0] == "end":
			paths[parts[1]] = append(paths[parts[1]], parts[0])

		case parts[0] == "start":
			paths[parts[0]] = append(paths[parts[0]], parts[1])

		case parts[1] == "end":
			paths[parts[0]] = append(paths[parts[0]], parts[1])

		case parts[1] == "start":
			paths[parts[1]] = append(paths[parts[1]], parts[0])

		default:
			paths[parts[0]] = append(paths[parts[0]], parts[1])
			paths[parts[1]] = append(paths[parts[1]], parts[0])
		}
	}

	result1 := traverse(paths, "start", make(map[string]bool))
	fmt.Println("Part 1: ", len(result1))

	result2 := traverseWithDoubleVisit(paths, "start", make(map[string]bool), false)
	fmt.Println("Part 2: ", len(result2))
}
