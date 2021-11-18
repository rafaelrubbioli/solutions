package main

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] < satisfaction[j]
	})

	result := 0
	for i := range satisfaction {
		sum := 0
		for _, tail := range satisfaction[i:] {
			sum += tail
		}

		if sum > 0 {
			result += sum
		}
	}

	return result
}

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
	fmt.Println(maxSatisfaction([]int{-2, 5, -1, 0, 3, -3}))
}
