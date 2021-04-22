package solution

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	counters := make(map[int]int)

	maxValue := 0
	lastMaxCounterValue := 0

	for _, x := range A {
		if x <= N {
			_, ok := counters[x]

			if !ok {
				counters[x] = lastMaxCounterValue
			}

			counters[x]++

			if maxValue < counters[x] {
				maxValue = counters[x]
			}
		}
		if x == N+1 {
			lastMaxCounterValue = maxValue
			counters = map[int]int{}
		}
	}

	result := make([]int, N)
	for i := 0; i < N; i++ {
		if val, ok := counters[i+1]; ok {
			result[i] = val
		} else {
			result[i] = lastMaxCounterValue
		}
	}

	return result
}
