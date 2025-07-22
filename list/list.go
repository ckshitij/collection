package list

import (
	"errors"
	"sync"
)

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
	mu   sync.RWMutex
}

var (
	ErrInvalidPosition  = errors.New("invalid position, please check the list size")
	ErrNegativePosition = errors.New("position must be non-negative")
	ErrOutOfBound       = errors.New("position out of bounds")
)

// Create a new list with the initializer element
func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// insert element in the beginning of the list
func (list *List[T]) PushFront(element T) {
	list.mu.Lock()
	defer list.mu.Unlock()

	newNode := NewNode(element)
	if list.head != nil {
		newNode.next = list.head
		list.head.prev = newNode
	} else {
		list.tail = newNode
	}
	list.head = newNode
	list.size++
}

// insert element in the end of the list
func (list *List[T]) PushBack(element T) {
	list.mu.Lock()
	defer list.mu.Unlock()

	newNode := NewNode(element)
	if list.tail != nil {
		newNode.prev = list.tail
		list.tail.next = newNode
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.size++
}

// Remove element from the beginning of the list
func (list *List[T]) PopFront() {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head == nil {
		return
	}
	list.head = list.head.Next()
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.size--
}

// Remove element from the end of the list
func (list *List[T]) PopBack() {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head == nil {
		return
	}
	list.tail = list.tail.Prev()
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = (nil)
	}
	list.size--
}

// Get element from the beginning of the list
func (list *List[T]) Front() *Node[T] {
	list.mu.RLock()
	defer list.mu.RUnlock()

	return list.head
}

// Get element from the back of the list
func (list *List[T]) Back() *Node[T] {
	list.mu.RLock()
	defer list.mu.RUnlock()

	return list.tail
}

// Return the size of list
func (list *List[T]) Len() int {
	list.mu.RLock()
	defer list.mu.RUnlock()

	return list.size
}

// InsertAtPosition inserts a node at a specific position in the list
func (list *List[T]) InsertAtPosition(data T, position int) error {
	list.mu.Lock()
	defer list.mu.Unlock()

	if position < 0 {
		return ErrNegativePosition
	}

	newNode := NewNode(data)
	if position == 0 {
		list.pushFrontNode(newNode)
		return nil
	}

	current := list.head
	for i := 0; i < position-1; i++ {
		if current == nil {
			return ErrOutOfBound
		}
		current = current.Next()
	}
	if current == nil || current.Next() == nil {
		list.pushBackNode(newNode)
	} else {
		newNode.next = (current.Next())
		newNode.prev = (current)
		current.Next().prev = (newNode)
		current.next = (newNode)
		list.size++
	}
	return nil
}

func (list *List[T]) IterateForward(action func(index int, element T)) {
	list.mu.RLock()
	defer list.mu.RUnlock()

	itr := list.head
	index := 0
	for itr != nil {
		action(index, itr.Element())
		itr = itr.Next()
		index++
	}
}

// Internal methods (non-locked, used within locked context)
func (list *List[T]) pushFrontNode(node *Node[T]) {
	if list.head != nil {
		node.next = (list.head)
		list.head.prev = (node)
	} else {
		list.tail = node
	}
	list.head = node
	list.size++
}

func (list *List[T]) pushBackNode(node *Node[T]) {
	if list.tail != nil {
		node.prev = (list.tail)
		list.tail.next = (node)
	} else {
		list.head = node
	}
	list.tail = node
	list.size++
}

func (list *List[T]) IterateBackward(action func(index int, element T)) {
	list.mu.RLock()
	defer list.mu.RUnlock()

	itr := list.tail
	index := list.size - 1
	for itr != nil {
		action(index, itr.Element())
		itr = itr.Prev()
		index--
	}
}

func (list *List[T]) Clear() {
	list.mu.Lock()
	defer list.mu.Unlock()

	list.head = nil
	list.tail = nil
	list.size = 0
}
