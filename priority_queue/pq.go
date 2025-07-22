package pq

import (
	"sync"
)

// Comparable defines a function type for comparing two elements of type T.
type Comparable[T any] func(a T, b T) bool

// PriorityQueue represents a thread-safe generic priority queue backed by a heap.
type PriorityQueue[T any] struct {
	table   []T
	compare Comparable[T]
	mu      sync.RWMutex
}

// NewPriorityQueue initializes a new thread-safe PriorityQueue.
func NewPriorityQueue[T any](compFunc Comparable[T], elements ...T) *PriorityQueue[T] {
	q := &PriorityQueue[T]{
		table:   append([]T{}, elements...),
		compare: compFunc,
	}
	q.buildHeap()
	return q
}

// Push inserts a new element while maintaining the heap property.
func (pq *PriorityQueue[T]) Push(data T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.table = append(pq.table, data)
	pq.siftUp(len(pq.table) - 1)
}

// Pop removes and returns the element with the highest priority.
func (pq *PriorityQueue[T]) Pop() T {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	var zero T
	if len(pq.table) == 0 {
		return zero
	}
	top := pq.table[0]
	lastIndex := len(pq.table) - 1
	pq.swap(0, lastIndex)
	pq.table = pq.table[:lastIndex]
	pq.heapify(0)
	return top
}

// Peek returns the highest-priority element without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	var zero T
	if len(pq.table) == 0 {
		return zero, false
	}
	return pq.table[0], true
}

// Size returns the number of elements.
func (pq *PriorityQueue[T]) Size() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	return len(pq.table)
}

// Empty returns true if the queue is empty.
func (pq *PriorityQueue[T]) Empty() bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	return len(pq.table) == 0
}

// Clear removes all elements.
func (pq *PriorityQueue[T]) Clear() {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.table = nil
}

// GetValues returns a copy slice between start and end indices (inclusive).
func (pq *PriorityQueue[T]) GetValues(startInd, endInd int) []T {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if startInd < 0 || endInd >= len(pq.table) || startInd > endInd {
		return []T{}
	}
	tempData := make([]T, endInd-startInd+1)
	copy(tempData, pq.table[startInd:endInd+1])
	return tempData
}

// --- Private methods (assume caller has lock) ---

func (_ *PriorityQueue[T]) parent(index int) int {
	return (index - 1) / 2
}

func (_ *PriorityQueue[T]) left(index int) int {
	return 2*index + 1
}

func (_ *PriorityQueue[T]) right(index int) int {
	return 2*index + 2
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.table[i], pq.table[j] = pq.table[j], pq.table[i]
}

func (pq *PriorityQueue[T]) heapify(index int) {
	highest := index
	left := pq.left(index)
	right := pq.right(index)

	if left < len(pq.table) && pq.compare(pq.table[left], pq.table[highest]) {
		highest = left
	}
	if right < len(pq.table) && pq.compare(pq.table[right], pq.table[highest]) {
		highest = right
	}

	if highest != index {
		pq.swap(index, highest)
		pq.heapify(highest)
	}
}

func (pq *PriorityQueue[T]) siftUp(index int) {
	for index > 0 {
		parent := pq.parent(index)
		if pq.compare(pq.table[index], pq.table[parent]) {
			pq.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

func (pq *PriorityQueue[T]) buildHeap() {
	for i := len(pq.table)/2 - 1; i >= 0; i-- {
		pq.heapify(i)
	}
}
