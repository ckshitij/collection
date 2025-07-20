package pq

// Comparable defines a function type for comparing two elements of type T.
// It should return true if the first argument has higher priority than the second.
// This allows customizing the behavior of the priority queue (e.g., min-heap, max-heap).
type Comparable[T any] func(a T, b T) bool

// PriorityQueue represents a generic priority queue backed by a heap.
// It stores elements of type T and uses the provided comparison function to maintain priority order.
type PriorityQueue[T any] struct {
	table   []T           // The internal slice holding the elements.
	compare Comparable[T] // The comparison function defining element priority.
}

// NewPriorityQueue initializes a new PriorityQueue with the given comparison function.
// Optionally, initial elements can be passed to populate the queue.
// The queue is heapified to ensure the heap property is satisfied.
func NewPriorityQueue[T any](compFunc Comparable[T], elements ...T) *PriorityQueue[T] {
	q := &PriorityQueue[T]{
		table:   append([]T{}, elements...),
		compare: compFunc,
	}
	q.buildHeap()
	return q
}

// Push inserts a new element into the priority queue, maintaining the heap property.
func (pq *PriorityQueue[T]) Push(data T) {
	pq.table = append(pq.table, data)
	pq.siftUp(len(pq.table) - 1)
}

// Pop removes and returns the element with the highest priority from the queue.
// If the queue is empty, it returns the zero value of T.
func (pq *PriorityQueue[T]) Pop() T {
	var zero T
	if pq.Empty() {
		return zero
	}
	top := pq.table[0]
	lastIndex := pq.Size() - 1
	pq.swap(0, lastIndex)
	pq.table = pq.table[:lastIndex]
	pq.heapify(0)
	return top
}

// Peek returns the element with the highest priority without removing it.
// It returns the zero value of T and false if the queue is empty.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	var zero T
	if pq.Empty() {
		return zero, false
	}
	return pq.table[0], true
}

// Size returns the number of elements currently in the priority queue.
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.table)
}

// Empty returns true if the priority queue contains no elements.
func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.table) == 0
}

// Clear removes all elements from the priority queue, leaving it empty.
func (pq *PriorityQueue[T]) Clear() {
	pq.table = nil
}

// GetValues returns a slice of elements from the priority queue between the given start and end indices (inclusive).
// If the indices are invalid, it returns an empty slice.
func (pq *PriorityQueue[T]) GetValues(startInd, endInd int) []T {
	if startInd < 0 || endInd >= pq.Size() || startInd > endInd {
		return []T{}
	}
	tempData := make([]T, endInd-startInd+1)
	copy(tempData, pq.table[startInd:endInd+1])
	return tempData
}

// parent returns the index of the parent node of the given index in the heap.
func (_ *PriorityQueue[T]) parent(index int) int {
	return (index - 1) / 2
}

// left returns the index of the left child of the given index in the heap.
func (_ *PriorityQueue[T]) left(index int) int {
	return 2*index + 1
}

// right returns the index of the right child of the given index in the heap.
func (_ *PriorityQueue[T]) right(index int) int {
	return 2*index + 2
}

// swap exchanges the elements at the two specified indices.
func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.table[i], pq.table[j] = pq.table[j], pq.table[i]
}

// heapify restores the heap property starting from the given index downwards.
func (pq *PriorityQueue[T]) heapify(index int) {
	highest := index
	left := pq.left(index)
	right := pq.right(index)

	if left < pq.Size() && pq.compare(pq.table[left], pq.table[highest]) {
		highest = left
	}
	if right < pq.Size() && pq.compare(pq.table[right], pq.table[highest]) {
		highest = right
	}

	if highest != index {
		pq.swap(index, highest)
		pq.heapify(highest)
	}
}

// siftUp restores the heap property by moving the element at the given index up the tree.
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

// buildHeap transforms the current elements into a valid heap in-place.
func (pq *PriorityQueue[T]) buildHeap() {
	for i := pq.Size()/2 - 1; i >= 0; i-- {
		pq.heapify(i)
	}
}
