package queue

import (
	"errors"

	"github.com/ckshitij/collection/list"
)

// Queue represents a generic queue data structure that holds elements of any type.
type Queue[T any] struct {
	head *list.List[T]
}

// NewQueue creates and returns a new instance of Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: list.NewList[T](),
	}
}

// Enqueue adds a new element to the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.head.PushBack(value)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *Queue[T]) IsEmpty() bool {
	return q.head.Len() == 0
}

// Dequeue removes the front element from the queue.
// Returns an error if the queue is empty.
func (q *Queue[T]) Dequeue() error {
	if q.IsEmpty() {
		return errors.New("invalid operation: empty queue")
	}
	q.head.PopFront()
	return nil
}

// Front returns the front element of the queue without removing it.
// Returns the zero value of the type if the queue is empty.
func (q *Queue[T]) Front() T {
	var zero T
	if node := q.head.Front(); node != nil {
		return node.Element()
	}
	return zero
}

// Back returns the back element of the queue without removing it.
// Returns the zero value of the type if the queue is empty.
func (q *Queue[T]) Back() T {
	var zero T
	if node := q.head.Back(); node != nil {
		return node.Element()
	}
	return zero
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.head.Len()
}
