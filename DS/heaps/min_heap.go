package min_heap

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type MinHeap struct {
	Items []int
	size  int
}

func (h *MinHeap) hasParent(idx int) (int, bool) {
	parentIdx := int(math.Ceil((float64(idx) - 2) / 2))

	return parentIdx, parentIdx >= 0
}

func (h *MinHeap) parent(idx int) (parent int, err error) {
	parentIdx, ok := h.hasParent(idx)

	if ok {
		parent = h.Items[parentIdx]
	} else {
		err = errors.New("no parent found, already at root")
	}

	return
}

func (h *MinHeap) hasLeftChild(idx int) (int, bool) {
	leftChildIdx := (idx * 2) + 1

	return leftChildIdx, leftChildIdx < h.size
}

func (h *MinHeap) leftChild(idx int) (leftChild int, err error) {
	leftChildIdx, ok := h.hasLeftChild(idx)

	if ok {
		leftChild = h.Items[leftChildIdx]
	} else {
		err = errors.New("no left child found")
	}

	return
}

func (h *MinHeap) hasRightChild(idx int) (int, bool) {
	rightChildIdx := (idx * 2) + 2

	return rightChildIdx, rightChildIdx < h.size
}

func (h *MinHeap) rightChild(idx int) (rightChild int, err error) {
	rightChildIdx, ok := h.hasRightChild(idx)

	if ok {
		rightChild = h.Items[rightChildIdx]
	} else {
		err = errors.New("no right child found")
	}

	return
}

func (h *MinHeap) swap(idx, otherIdx int) {
	h.Items[idx], h.Items[otherIdx] = h.Items[otherIdx], h.Items[idx]
}

func (h *MinHeap) heapifyDown() {
	currIdx := 0

	for {
		leftIdx, ok := h.hasLeftChild(currIdx)

		if !ok {
			return
		}

		smallerChildIdx := leftIdx
		if rightIdx, ok := h.hasRightChild(currIdx); ok && h.Item(rightIdx) < h.Item(leftIdx) {
			smallerChildIdx = rightIdx
		}

		if h.Item(currIdx) < h.Item(smallerChildIdx) {
			break
		}

		h.swap(currIdx, smallerChildIdx)
		currIdx = smallerChildIdx
	}
}

func (h *MinHeap) Peek() (int, error) {
	if h.size == 0 {
		return 0, errors.New("heap has no items")
	}

	return h.Item(0), nil
}

func (h *MinHeap) Poll() (int, error) {
	if h.size == 0 {
		return 0, errors.New("heap has no items")
	}

	item := h.Items[0]

	h.Items[0] = h.Items[h.size-1]

	h.size--

	h.Items = h.Items[0:h.size] // Resizing the slice

	h.heapifyDown()

	return item, nil
}

func (h *MinHeap) Item(idx int) int {
	return h.Items[idx]
}

func (h *MinHeap) heapifyUp() {
	currIdx := h.size - 1

	for {
		parentIdx, ok := h.hasParent(currIdx)

		if !ok || h.Item(parentIdx) < h.Item(currIdx) {
			return
		}
		h.swap(currIdx, parentIdx)
		currIdx = parentIdx
	}
}

func (h *MinHeap) Add(newItem int) {
	h.Items = append(h.Items, newItem)
	h.size++

	h.heapifyUp()
}

func (h MinHeap) String() string {
	output := []string{""}
	currIdx := 0
	currLevel := 0
	itemsCanBeFit := 0

	for currIdx < h.size {
		output[currLevel] += fmt.Sprintf("%d\t", h.Item(currIdx))
		// fmt.Printf("Curr Level: %#v | Curr Idx: %#v | current output: %#v | Final Output: %#v\n", currLevel, currIdx, output[currLevel], output)
		currIdx++

		if currIdx > itemsCanBeFit {
			currLevel++
			itemsCanBeFit += int(math.Pow(2, float64(currLevel)))
			output = append(output, "")
		}
	}

	return fmt.Sprintf("Size: %d | Items: %#v\nTree:\n%s\n------------------------------\n", h.size, h.Items, strings.Join(output, "\n"))
}
