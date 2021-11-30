package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	say := countAndSay(n - 1)

	fmt.Println(say)
	count := 1
	result := ""
	lastSeen := ""

	for _, v := range say {
		char := string(v)
		if lastSeen == "" {
			lastSeen = char
			continue
		}

		if char == lastSeen {
			count++
			continue
		}

		result += strconv.Itoa(count) + lastSeen
		count = 1
		lastSeen = char
	}

	result += strconv.Itoa(count) + lastSeen

	return result
}

func main() {
	fmt.Println(countAndSay(4))
}
