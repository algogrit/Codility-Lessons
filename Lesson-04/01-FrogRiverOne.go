package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	count := 0
	hasFallenLeaf := make(map[int]bool)
	for idx, el := range A {
		if el <= X {
			hasLeafInPos := hasFallenLeaf[el]

			if !hasLeafInPos {
				count++
				hasFallenLeaf[el] = true
			}
		}
		if count == X {
			return idx
		}
	}

	return -1
}
