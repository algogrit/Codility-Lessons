package solution

import "math"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sum := 0
	for _, el := range A {
		sum += el
	}

	minDistance := math.Inf(1)
	runningSum := 0
	for idx, el := range A {
		if idx+1 == len(A) {
			break
		}

		runningSum += el
		distance := math.Abs(float64((sum - runningSum) - runningSum))

		if distance < minDistance {
			minDistance = distance
		}
	}

	return int(minDistance)
}
