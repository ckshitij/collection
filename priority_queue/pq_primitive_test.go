package pq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinIntPQ(t *testing.T) {
	pq := NewMinIntPQ(4, 1, 3, 2)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxIntPQ(t *testing.T) {
	pq := NewMaxIntPQ(4, 1, 3, 2)
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinInt64PQ(t *testing.T) {
	pq := NewMinInt64PQ(4, 1, 3, 2)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxInt64PQ(t *testing.T) {
	pq := NewMaxInt64PQ(4, 1, 3, 2)
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinInt32PQ(t *testing.T) {
	pq := NewMinInt32PQ(4, 1, 3, 2)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxInt32PQ(t *testing.T) {
	pq := NewMaxInt32PQ(4, 1, 3, 2)
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinInt16PQ(t *testing.T) {
	pq := NewMinInt16PQ(4, 1, 3, 2)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxInt16PQ(t *testing.T) {
	pq := NewMaxInt16PQ(4, 1, 3, 2)
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinInt8PQ(t *testing.T) {
	pq := NewMinInt8PQ(4, 1, 3, 2)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxInt8PQ(t *testing.T) {
	pq := NewMaxIntPQ(4, 1, 3, 2)
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 2, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinFloat64PQ(t *testing.T) {
	pq := NewMinFloat64PQ(3.5, 2.1, 4.9, 1.2)
	assert.Equal(t, 1.2, pq.Pop())
	assert.Equal(t, 2.1, pq.Pop())
	assert.Equal(t, 3.5, pq.Pop())
	assert.Equal(t, 4.9, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxFloat64PQ(t *testing.T) {
	pq := NewMaxFloat64PQ(3.5, 2.1, 4.9, 1.2)
	assert.Equal(t, 4.9, pq.Pop())
	assert.Equal(t, 3.5, pq.Pop())
	assert.Equal(t, 2.1, pq.Pop())
	assert.Equal(t, 1.2, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinFloat32PQ(t *testing.T) {
	pq := NewMinFloat32PQ(3.5, 2.1, 4.9, 1.2)
	assert.Equal(t, 1.2, pq.Pop())
	assert.Equal(t, 2.1, pq.Pop())
	assert.Equal(t, 3.5, pq.Pop())
	assert.Equal(t, 4.9, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxFloat32PQ(t *testing.T) {
	pq := NewMaxFloat64PQ(3.5, 2.1, 4.9, 1.2)
	assert.Equal(t, 4.9, pq.Pop())
	assert.Equal(t, 3.5, pq.Pop())
	assert.Equal(t, 2.1, pq.Pop())
	assert.Equal(t, 1.2, pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMinStringPQ(t *testing.T) {
	pq := NewMinStringPQ("delta", "alpha", "charlie", "bravo")
	assert.Equal(t, "alpha", pq.Pop())
	assert.Equal(t, "bravo", pq.Pop())
	assert.Equal(t, "charlie", pq.Pop())
	assert.Equal(t, "delta", pq.Pop())
	assert.True(t, pq.Empty())
}

func TestNewMaxStringPQ(t *testing.T) {
	pq := NewMaxStringPQ("delta", "alpha", "charlie", "bravo")
	assert.Equal(t, "delta", pq.Pop())
	assert.Equal(t, "charlie", pq.Pop())
	assert.Equal(t, "bravo", pq.Pop())
	assert.Equal(t, "alpha", pq.Pop())
	assert.True(t, pq.Empty())
}
