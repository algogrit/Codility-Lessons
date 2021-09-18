package main

import "fmt"

func findCenter(edges [][]int) int {
	var firstNode, secondNode int

	firstNode = edges[0][0]
	secondNode = edges[0][1]

	for _, edge := range edges[1:2] {
		if edge[0] == firstNode || edge[1] == firstNode {
			return firstNode
		}
	}

	return secondNode
}

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))

	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}
