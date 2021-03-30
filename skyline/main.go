package main

import "fmt"

func main() {
	heights := make(map[int]int)
	var last int
	for {
		var x, y, height int
		if count, _ := fmt.Scanf("%d %d %d", &x, &height, &y); count != 3 {
			break
		}

		last = y
		for i := x; i < y; i++ {
			if h, ok := heights[i]; ok {
				if height <= h {
					continue
				}
			}

			heights[i] = height
		}

	}

	result := make([]int, last+1)
	for i, r := range heights {
		result[i-1] = r
	}

	current := 0
	first := false
	for i, r := range result {
		if !first {
			fmt.Print(i+1, " ", r)
			current = r
			first = true
			continue
		}

		if current != r {
			fmt.Print(" ", i+1, " ", r)
			current = r
		}
	}
	fmt.Println("")
}
