package list

import (
	"errors"
)

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size uint
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
	newNode := NewNode(element)
	if list.head != nil {
		newNode.SetNext(list.head)
		list.head.SetPrevious(newNode)
	} else {
		list.tail = newNode
	}
	list.head = newNode
	list.size++
}

// insert element in the end of the list
func (list *List[T]) PushBack(element T) {
	newNode := NewNode(element)
	if list.tail != nil {
		newNode.SetPrevious(list.tail)
		list.tail.SetNext(newNode)
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.size++
}

// Remove element from the beginning of the list
func (list *List[T]) PopFront() {
	if list.head == nil {
		return
	}
	list.head = list.head.Next()
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.SetPrevious(nil)
	}
	list.size--
}

// Remove element from the end of the list
func (list *List[T]) PopBack() {
	if list.head == nil {
		return
	}
	list.tail = list.tail.Prev()
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.SetNext(nil)
	}
	list.size--
}

// Get element from the beginning of the list
func (list *List[T]) Front() *Node[T] {
	return list.head
}

// Get element from the back of the list
func (list *List[T]) Back() *Node[T] {
	return list.tail
}

// Return the size of list
func (list *List[T]) Len() uint {
	return list.size
}

// InsertAtPosition inserts a node at a specific position in the list
func (dll *List[T]) InsertAtPosition(data T, position int) error {
	newNode := NewNode(data)

	if position < 0 {
		return ErrNegativePosition
	}

	if position == 0 {
		dll.PushFront(data)
		return nil
	}

	current := dll.head
	for i := 0; i < position-1; i++ {
		if current == nil {
			return ErrOutOfBound
		}
		current = current.Next()
	}

	if current == nil || current.Next() == nil {
		dll.PushBack(data)
	} else {
		newNode.SetNext(current.Next())
		newNode.SetPrevious(current)
		current.Next().SetPrevious(newNode)
		current.SetNext(newNode)
		dll.size++
	}
	return nil
}

func (list *List[T]) IterateForward(action func(index int, element T)) {
	if list.head == nil {
		return // No-op for an empty list
	}
	itr := list.head
	i := 0
	for itr != nil {
		action(i, itr.Element())
		itr = itr.Next()
		i++
	}
}

func (list *List[T]) IterateBackward(action func(index uint, element T)) {
	if list.tail == nil {
		return // No-op for an empty list
	}
	itr := list.tail
	i := list.size - 1
	for itr != nil {
		action(i, itr.Element())
		itr = itr.Prev()
		i--
	}
}

func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}
