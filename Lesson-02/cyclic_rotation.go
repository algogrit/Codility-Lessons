package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	sizeOfArr := len(A)

	if sizeOfArr == 0 {
		return nil
	}

	K = K % sizeOfArr

	tail := A[:sizeOfArr-K]
	head := A[sizeOfArr-K:]

	result := make([]int, len(head), sizeOfArr)

	copy(result, head)

	for _, el := range tail {
		result = append(result, el)
	}

	return result
}
