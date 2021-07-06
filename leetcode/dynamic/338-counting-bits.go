func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}

	return res
}
