package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type entry struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func overlappingPerpendicularLines(entries []entry) int {
	seen := make(map[string]int)
	for _, entry := range entries {
		if entry.x1 == entry.x2 {
			if entry.y1 > entry.y2 {
				entry.y1, entry.y2 = entry.y2, entry.y1
			}

			for y := entry.y1; y <= entry.y2; y++ {
				seen[key(entry.x1, y)]++
			}
		} else if entry.y1 == entry.y2 {
			if entry.x1 > entry.x2 {
				entry.x1, entry.x2 = entry.x2, entry.x1
			}

			for x := entry.x1; x <= entry.x2; x++ {
				seen[key(x, entry.y1)]++
			}
		}
	}

	overlapping := 0
	for _, v := range seen {
		if v > 1 {
			overlapping++
		}
	}

	return overlapping
}

func getLine(e entry) func(int) int {
	m := (e.y2 - e.y1) / (e.x2 - e.x1)
	b := e.y1 - m*e.x1
	return func(i int) int {
		return m*i + b
	}
}

func overlappingLines(entries []entry) int {
	seen := make(map[string]int)
	for _, entry := range entries {
		if entry.x1 == entry.x2 {
			if entry.y1 > entry.y2 {
				entry.y1, entry.y2 = entry.y2, entry.y1
			}

			for y := entry.y1; y <= entry.y2; y++ {
				seen[key(entry.x1, y)]++
			}
		} else if entry.y1 == entry.y2 {
			if entry.x1 > entry.x2 {
				entry.x1, entry.x2 = entry.x2, entry.x1
			}

			for x := entry.x1; x <= entry.x2; x++ {
				seen[key(x, entry.y1)]++
			}
		} else {
			f := getLine(entry)
			if entry.x1 > entry.x2 {
				entry.x1, entry.x2 = entry.x2, entry.x1
			}

			for x := entry.x1; x <= entry.x2; x++ {
				seen[key(x, f(x))]++
			}
		}
	}

	overlapping := 0
	for _, v := range seen {
		if v > 1 {
			overlapping++
		}
	}

	return overlapping
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	entries := make([]entry, 0)
	for scanner.Scan() {
		var x1, x2, y1, y2 int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}

		entries = append(entries, entry{x1: x1, y1: y1, x2: x2, y2: y2})
	}

	fmt.Println("part 1: ", overlappingPerpendicularLines(entries))
	fmt.Println("part 2: ", overlappingLines(entries))
}
