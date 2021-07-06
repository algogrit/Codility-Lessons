func hammingDistance(x int, y int) int {
	if x == y {
		return 0
	}

	xIsOdd := x % 2
	yIsOdd := y % 2

	if xIsOdd != yIsOdd {
		return 1 + hammingDistance(x>>1, y>>1)
	} else {
		return hammingDistance(x>>1, y>>1)
	}
}
