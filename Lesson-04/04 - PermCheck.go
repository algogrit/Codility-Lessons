package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	n := len(A)
	sum := n * (n + 1) / 2

	isSeen := make(map[int]bool)

	for _, el := range A {
		dup := isSeen[el]
		if dup {
			return 0
		}
		isSeen[el] = true
		sum -= el
	}
	if sum != 0 {
		return 0
	}
	return 1
}
