package solution

// Worst Case O(n * m)

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	counters := make([]int, N)

	maxValue := 0

	for _, x := range A {
		idx := x - 1
		if x <= N {
			counters[idx]++

			if maxValue < counters[idx] {
				maxValue = counters[idx]
			}
		}
		if x == N+1 {
			for idx := range counters {
				counters[idx] = maxValue
			}
		}
	}
	return counters
}
