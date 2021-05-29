package main

import "fmt"

func maxSubArray(nums []int) int {
	var currentSum int
	var maxSum = nums[0]

	for _, num := range nums {
		currentSum += num
		if maxSum < currentSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}

func main() {
	test1 := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(test1)

	test2 := maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(test2)

	test3 := maxSubArray([]int{-1})
	fmt.Println(test3)
}
