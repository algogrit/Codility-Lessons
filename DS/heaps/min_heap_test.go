package min_heap_test

import (
	"testing"

	mh "algogrit.com/ds/heaps"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	minHeap := &mh.MinHeap{}

	minHeap.Add(4)
	minHeap.Add(5)

	minHeap.Add(3)
	minHeap.Add(2)

	assert.Equal(t, []int{2, 3, 4, 5}, minHeap.Items)

	minHeap.Add(1)

	assert.Equal(t, []int{1, 2, 4, 5, 3}, minHeap.Items)
}

func TestPoll(t *testing.T) {
	minHeap := &mh.MinHeap{}

	minHeap.Add(4)
	minHeap.Add(5)

	minHeap.Add(3)
	minHeap.Add(2)

	assert.Equal(t, []int{2, 3, 4, 5}, minHeap.Items)

	minHeap.Add(1)

	assert.Equal(t, []int{1, 2, 4, 5, 3}, minHeap.Items)

	el, err := minHeap.Poll()
	assert.Equal(t, 1, el)
	assert.Nil(t, err)

	el, err = minHeap.Poll()
	assert.Equal(t, 2, el)
	assert.Nil(t, err)

	el, err = minHeap.Poll()
	assert.Equal(t, 3, el)
	assert.Nil(t, err)

	el, err = minHeap.Poll()
	assert.Equal(t, 4, el)
	assert.Nil(t, err)

	el, err = minHeap.Poll()
	assert.Equal(t, 5, el)
	assert.Nil(t, err)

	_, err = minHeap.Poll()
	assert.NotNil(t, err)
}
