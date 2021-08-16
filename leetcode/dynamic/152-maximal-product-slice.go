package main

import "fmt"

// given an integer array. find the product of the maximum product subarray

// {6, -3, -10, 0, 2}
// 180

// {-1, -3, -10, 0, 60}
// 60

// -100 < el < 100

// [-1] => 0 => [] => 0

func maxInt(nums ...int) int {
	maxNum := nums[0]

	for _, el := range nums {
		if el > maxNum {
			maxNum = el
		}
	}

	return maxNum
}

func minInt(nums ...int) int {
	minNum := nums[0]

	for _, el := range nums {
		if el < minNum {
			minNum = el
		}
	}

	return minNum
}

func maximalSubProduct(arr []int) int {
	maxProduct := 0

	currMax := 1
	currMin := 1

	for _, num := range arr {
		if num == 0 {
			currMax = 1
			currMin = 1

			continue
		}

		// prevMax := currMax
		currMax, currMin = maxInt(currMax*num, currMin*num, num), minInt(currMax*num, currMin*num, num)

		maxProduct = maxInt(currMax, maxProduct)
	}

	return maxProduct
}

func main() {
	arr := []int{6, -3, -10, 0, 2}

	fmt.Println(maximalSubProduct(arr))
}
