package list

type Node[T any] struct {
	element T
	prev    *Node[T]
	next    *Node[T]
}

func NewNode[T any](element T) *Node[T] {
	return &Node[T]{element: element}
}

func (n *Node[T]) Element() T {
	return n.element
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}
