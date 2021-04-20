package solution

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	binStr := fmt.Sprintf("%b", N)
	maxCountOfZeros := 0
	countOfZeros := 0
	isInZero := false
	hasSeenOne := false

	for _, el := range binStr {
		if isInZero {
			if string(el) == "0" {
				countOfZeros++
			} else {
				if maxCountOfZeros < countOfZeros {
					maxCountOfZeros = countOfZeros
				}
				isInZero = false
			}
		} else {
			if string(el) == "0" {
				if hasSeenOne {
					isInZero = true
					countOfZeros = 1
				}
			} else {
				hasSeenOne = true
			}
		}
	}
	return maxCountOfZeros
}
