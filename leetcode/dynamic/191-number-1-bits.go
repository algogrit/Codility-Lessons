var memo = map[uint32]int{}

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	if res, ok := memo[num]; ok {
		return res
	}

	res := hammingWeight(num >> 1)

	if num%2 == 1 {
		res = res + 1
	}

	memo[num] = res

	return res
}

// func hammingWeight(num uint32) int {
// 	if num == 0 {
// 		return 0
// 	}

// 	if num%2 == 1 {
// 		return 1 + hammingWeight(num>>1)
// 	} else {
// 		return hammingWeight(num >> 1)
// 	}
// }
