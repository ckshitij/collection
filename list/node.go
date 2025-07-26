package list

import "sync"

// Node represents a thread-safe element in a doubly linked list.
type Node[T any] struct {
	element T
	prev    *Node[T]
	next    *Node[T]
	mu      sync.RWMutex
}

// NewNode creates a new node with the given element.
func NewNode[T any](element T) *Node[T] {
	return &Node[T]{element: element}
}

func (n *Node[T]) Element() T {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.element
}

func (n *Node[T]) Next() *Node[T] {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.prev
}

// internal helpers for safe mutation (called only under list lock)
func (n *Node[T]) setNext(next *Node[T]) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.next = next
}

func (n *Node[T]) setPrev(prev *Node[T]) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.prev = prev
}
