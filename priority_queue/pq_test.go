package pq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriorityQueue(t *testing.T) {
	// Min-heap comparator
	minHeapComp := func(a, b int) bool {
		return a < b
	}

	// Test creating an empty queue
	pq := NewPriorityQueue(minHeapComp)
	assert.Equal(t, 0, pq.Size(), "Priority queue size should be zero on creation")

	// Test creating a queue with initial elements
	pq = NewPriorityQueue(minHeapComp, 5, 3, 8, 1)
	assert.Equal(t, 4, pq.Size(), "Priority queue size should match the number of initial elements")
	assert.Equal(t, 1, pq.Pop(), "Queue should pop the smallest element in a min-heap")

	// Test creating a queue with duplicate elements
	pq = NewPriorityQueue(minHeapComp, 5, 5, 5)
	assert.Equal(t, 3, pq.Size(), "Priority queue size should match the number of initial elements")
	assert.Equal(t, 5, pq.Pop(), "Queue should handle duplicate elements correctly")
}

func TestPriorityQueuePush(t *testing.T) {
	maxHeapComp := func(a, b int) bool {
		return a > b
	}
	pq := NewPriorityQueue(maxHeapComp)

	pq.Push(10)
	pq.Push(20)
	pq.Push(5)

	assert.Equal(t, 3, pq.Size(), "Priority queue size should be 3 after pushes")
	assert.Equal(t, 20, pq.Pop(), "Should pop the largest element in a max-heap")

	pq.Push(1000000)
	pq.Push(-1000000)
	assert.Equal(t, 1000000, pq.Pop(), "Should handle very large numbers correctly")
	assert.Equal(t, 10, pq.Pop())
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, -1000000, pq.Pop())
}

func TestPriorityQueuePop(t *testing.T) {
	minHeapComp := func(a, b int) bool {
		return a < b
	}
	pq := NewPriorityQueue(minHeapComp, 15, 10, 30, 5)

	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 10, pq.Pop())
	assert.Equal(t, 15, pq.Pop())
	assert.Equal(t, 30, pq.Pop())

	assert.Equal(t, 0, pq.Size())
	var zeroVal int
	assert.Equal(t, zeroVal, pq.Pop(), "Popping from an empty queue should return zero value")
}

func TestPriorityQueueSize(t *testing.T) {
	pq := NewPriorityQueue(func(a, b int) bool { return a < b })

	assert.Equal(t, 0, pq.Size())

	pq.Push(1)
	pq.Push(2)
	assert.Equal(t, 2, pq.Size())

	pq.Pop()
	assert.Equal(t, 1, pq.Size())
}

func TestPriorityQueueEmpty(t *testing.T) {
	pq := NewPriorityQueue(func(a, b int) bool { return a < b })

	assert.True(t, pq.Empty())
	pq.Push(1)
	pq.Push(2)
	assert.False(t, pq.Empty())

	pq.Pop()
	pq.Pop()
	assert.True(t, pq.Empty())
}

func TestPriorityQueueValues(t *testing.T) {
	pq := NewPriorityQueue(func(a, b int) bool { return a < b }, 10, 25, 12, 8, 41)

	assert.Equal(t, []int{8, 10, 12}, pq.GetValues(0, 2))
	assert.Equal(t, []int{}, pq.GetValues(-1, 5))
	assert.Equal(t, pq.Size(), len(pq.GetValues(0, pq.Size()-1)))

	pq.Push(1)
	pq.Push(2)
	assert.Equal(t, []int{1}, pq.GetValues(0, 0))

	pq.Pop()
	pq.Pop()
	pq.Clear()
	assert.True(t, pq.Empty())
}

func TestPriorityQueueHeapify(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	pq := NewPriorityQueue(comp, 10, 20, 5)
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 10, pq.Pop())

	pq = NewPriorityQueue(comp)
	assert.Equal(t, 0, pq.Size())
}

func TestPriorityQueueSiftUp(t *testing.T) {
	minHeapComp := func(a, b int) bool {
		return a < b
	}
	pq := NewPriorityQueue(minHeapComp)

	pq.Push(20)
	pq.Push(10)

	assert.Equal(t, 10, pq.Pop())
	assert.Equal(t, 20, pq.Pop())

	pq.Push(50)
	assert.Equal(t, 50, pq.Pop())
}

func TestPriorityQueueCustomType(t *testing.T) {
	type customType struct {
		priority int
		name     string
	}

	comp := func(a, b customType) bool {
		return a.priority > b.priority
	}

	pq := NewPriorityQueue(comp)
	pq.Push(customType{priority: 3, name: "Alice"})
	pq.Push(customType{priority: 5, name: "Bob"})
	pq.Push(customType{priority: 1, name: "Charlie"})

	top := pq.Pop()
	assert.Equal(t, 5, top.priority)
	assert.Equal(t, "Bob", top.name)

	pq.Pop()
	pq.Pop()
	assert.Equal(t, 0, pq.Size())
}

func TestPriorityQueueEdgeCases(t *testing.T) {
	pq := NewPriorityQueue(func(a, b int) bool { return a < b })
	assert.Equal(t, 0, pq.Size())
	assert.Equal(t, 0, pq.Pop())

	pq.Push(42)
	assert.Equal(t, 1, pq.Size())
	assert.Equal(t, 42, pq.Pop())
	assert.Equal(t, 0, pq.Size())

	pq.Push(7)
	pq.Push(7)
	assert.Equal(t, 2, pq.Size())
	assert.Equal(t, 7, pq.Pop())
	assert.Equal(t, 7, pq.Pop())
	assert.Equal(t, 0, pq.Size())
}
