package main

import (
	"fmt"

	mh "algogrit.com/ds/heaps"
)

func main() {
	minHeap := &mh.MinHeap{}

	minHeap.Add(10)
	minHeap.Add(2)
	fmt.Println(minHeap)
	minHeap.Add(3)
	fmt.Println(minHeap)
	minHeap.Add(5)
	fmt.Println(minHeap)
	minHeap.Add(8)
	fmt.Println(minHeap)
	minHeap.Add(14)
	fmt.Println(minHeap)
	minHeap.Add(1)
	fmt.Println(minHeap)
	minHeap.Add(7)
	fmt.Println(minHeap)
}
