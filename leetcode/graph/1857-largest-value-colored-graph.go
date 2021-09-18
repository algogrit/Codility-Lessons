package main

import (
	"fmt"
	"strings"
)

type Node struct {
	ID            int
	Color         rune
	DirectedNodes []*Node
}

func constructNetwork(colors string, edges [][]int) ([]*Node, string) {
	network := make([]*Node, 0, len(colors))
	uniqColors := []rune{}

	for idx, color := range colors {
		node := new(Node)
		node.ID = idx
		node.Color = color
		node.DirectedNodes = make([]*Node, 0)

		if !strings.Contains(string(uniqColors), string(color)) {
			uniqColors = append(uniqColors, color)
		}

		network = append(network, node)
	}

	for _, edge := range edges {
		srcNode := network[edge[0]]
		destNode := network[edge[1]]

		srcNode.DirectedNodes = append(srcNode.DirectedNodes, destNode)
	}

	return network, string(uniqColors)
}

func contains(list []int, element int) bool {
	for _, el := range list {
		if element == el {
			return true
		}
	}
	return false
}

func max(nums ...int) int {
	maxNum := -100
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

type ColorCount map[rune]int

func colorCount(network []*Node, uniqColors string, memo map[int]ColorCount, visitedNodes []int, node *Node) ColorCount {
	if countMap, ok := memo[node.ID]; ok {
		return countMap
	}

	res := make(ColorCount)

	for _, directedNode := range node.DirectedNodes {
		if contains(visitedNodes, directedNode.ID) {
			panic("there is a cycle!")
		}

		countMap := colorCount(network, uniqColors, memo, append(visitedNodes, directedNode.ID), directedNode)

		for _, color := range uniqColors {
			count := countMap[color]
			res[color] = max(res[color], count)
		}
	}

	res[node.Color]++

	memo[node.ID] = res

	return res
}

func largestPathValue(colors string, edges [][]int) (maxColorCount int) {
	maxColorCount = -1
	defer func() {
		if r := recover(); r != nil {
			maxColorCount = -1
		}
	}()
	network, uniqColors := constructNetwork(colors, edges)

	memo := make(map[int]ColorCount)

	for _, node := range network {
		countMap := colorCount(network, uniqColors, memo, []int{node.ID}, node)

		for _, count := range countMap {
			if count > maxColorCount {
				maxColorCount = count
			}
		}
	}

	return
}

func main() {
	fmt.Println(largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}})) // 3

	fmt.Println(largestPathValue("a", [][]int{{0, 0}})) // -1
}
