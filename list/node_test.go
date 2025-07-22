package list

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(5)
	if node.Element() != 5 {
		t.Errorf("Expected element to be 5, got %v", node.Element())
	}
	if node.Next() != nil {
		t.Errorf("Expected next to be nil, got %v", node.Next())
	}
	if node.Prev() != nil {
		t.Errorf("Expected previous to be nil, got %v", node.Prev())
	}
}

func TestSetNext(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.next = node2

	if node1.Next() != node2 {
		t.Errorf("Expected node1's next to be node2, got %v", node1.Next())
	}
}

func TestSetPrevious(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node1.prev = node2

	if node1.Prev() != node2 {
		t.Errorf("Expected node1's previous to be node2, got %v", node1.Prev())
	}
}

func TestChainedNodes(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	if node1.Next() != node2 {
		t.Errorf("Expected node1's next to be node2, got %v", node1.Next())
	}

	if node2.Prev() != node1 {
		t.Errorf("Expected node2's previous to be node1, got %v", node2.Prev())
	}

	if node2.Next() != node3 {
		t.Errorf("Expected node2's next to be node3, got %v", node2.Next())
	}

	if node3.Prev() != node2 {
		t.Errorf("Expected node3's previous to be node2, got %v", node3.Prev())
	}
}
