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

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (list *List[T]) PushFront(element T) {
	list.mu.Lock()
	defer list.mu.Unlock()

	newNode := NewNode(element)
	if list.head != nil {
		newNode.setNext(list.head)
		list.head.setPrev(newNode)
	} else {
		list.tail = newNode
	}
	list.head = newNode
	list.size++
}

func (list *List[T]) PushBack(element T) {
	list.mu.Lock()
	defer list.mu.Unlock()

	newNode := NewNode(element)
	if list.tail != nil {
		newNode.setPrev(list.tail)
		list.tail.setNext(newNode)
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.size++
}

func (list *List[T]) PopFront() {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head == nil {
		return
	}
	next := list.head.Next()
	list.head = next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.setPrev(nil)
	}
	list.size--
}

func (list *List[T]) PopBack() {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.tail == nil {
		return
	}
	prev := list.tail.Prev()
	list.tail = prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.setNext(nil)
	}
	list.size--
}

func (list *List[T]) Front() *Node[T] {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.head
}

func (list *List[T]) Back() *Node[T] {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.tail
}

func (list *List[T]) Len() int {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.size
}

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
		next := current.Next()
		newNode.setNext(next)
		newNode.setPrev(current)
		next.setPrev(newNode)
		current.setNext(newNode)
		list.size++
	}
	return nil
}

func (list *List[T]) IterateForward(action func(index int, element T)) {
	list.mu.RLock()
	current := list.head
	index := 0
	list.mu.RUnlock()

	for current != nil {
		element := current.Element()
		next := current.Next()
		action(index, element)
		current = next
		index++
	}
}

func (list *List[T]) IterateBackward(action func(index int, element T)) {
	list.mu.RLock()
	current := list.tail
	index := list.size - 1
	list.mu.RUnlock()

	for current != nil {
		element := current.Element()
		prev := current.Prev()
		action(index, element)
		current = prev
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

// --- internal helpers (must be called under list.mu.Lock()) ---

func (list *List[T]) pushFrontNode(node *Node[T]) {
	if list.head != nil {
		node.setNext(list.head)
		list.head.setPrev(node)
	} else {
		list.tail = node
	}
	list.head = node
	list.size++
}

func (list *List[T]) pushBackNode(node *Node[T]) {
	if list.tail != nil {
		node.setPrev(list.tail)
		list.tail.setNext(node)
	} else {
		list.head = node
	}
	list.tail = node
	list.size++
}
