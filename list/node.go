package list

type Node[T any] struct {
	element  T
	previous *Node[T]
	next     *Node[T]
}

func NewNode[T any](element T) *Node[T] {
	return &Node[T]{
		element:  element,
		previous: nil,
		next:     nil,
	}
}

func (nn *Node[T]) Element() T {
	return nn.element
}

func (nn *Node[T]) Next() *Node[T] {
	return nn.next
}

func (nn *Node[T]) Prev() *Node[T] {
	return nn.previous
}

func (nn *Node[T]) SetNext(next *Node[T]) {
	nn.next = next
}

func (nn *Node[T]) SetPrevious(previous *Node[T]) {
	nn.previous = previous
}
