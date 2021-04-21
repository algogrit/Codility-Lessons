package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	n := len(A) + 1
	sum := float64(n*(n+1)) / 2

	for _, el := range A {
		sum -= float64(el)
	}

	return int(sum)
}
