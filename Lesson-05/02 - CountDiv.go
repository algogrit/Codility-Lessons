package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	res := (B - A) / K

	isADivByK := 0

	if A%K == 0 && A > 0 {
		isADivByK = 1
	}

	return res + isADivByK

	// // write your code in Go 1.4
	// countNDivByK := 0
	// for i := A; i <= B; i++ {
	// 	if i%K == 0 {
	// 		countNDivByK++
	// 	}
	// }

	// return countNDivByK
}
