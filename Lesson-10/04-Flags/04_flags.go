package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"fmt"
	"math"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	peaks := []int{}

	// for idx, el := range A {
	// 	if (idx+1 == len(A) || el > A[idx+1]) && (idx == 0 || el > A[idx-1]) {
	// 		peaks = append(peaks, idx)
	// 	}
	// }

	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) <= 1 {
		return len(peaks)
	}

	maxFlags := int(math.Ceil(math.Sqrt(float64(peaks[len(peaks)-1] - peaks[0]))))

	for flags := maxFlags; flags > 0; flags-- {
		startIdx := 0
		endIdx := len(peaks) - 1

		startFlag := peaks[startIdx]
		endFlag := peaks[endIdx]

		flagsPlaced := 2

		for startIdx < endIdx {
			startIdx++
			endIdx--

			if (peaks[startIdx] >= startFlag+flags) && (peaks[startIdx] <= endFlag-flags) {
				flagsPlaced++
				startFlag = peaks[startIdx]
			}

			if (peaks[endIdx] <= endFlag-flags) && (peaks[endIdx] >= startFlag+flags) {
				flagsPlaced++
				endFlag = peaks[endIdx]
			}

			if flagsPlaced >= flags {
				return flagsPlaced
			}
		}
	}

	return 0
}

func main() {
	// A := make([]int, 12)
	A := []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
	// A[0] = 1
	// A[1] = 5
	// A[2] = 3
	// A[3] = 4
	// A[4] = 3
	// A[5] = 4
	// A[6] = 1
	// A[7] = 2
	// A[8] = 3
	// A[9] = 4
	// A[10] = 6
	// A[11] = 2

	fmt.Println(Solution(A))
}
