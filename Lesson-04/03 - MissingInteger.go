package solution

import (
	"math"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	maxValue := math.MinInt64
	seenNumbers := make(map[int]bool)

	for _, el := range A {
		if maxValue < el {
			maxValue = el
		}

		if el > 0 {
			seenNumbers[el] = true
		}
	}

	if maxValue <= 0 {
		return 1
	}

	for i := 1; i <= maxValue; i++ {
		found := seenNumbers[i]
		if !found {
			return i
		}
	}

	return maxValue + 1
}
