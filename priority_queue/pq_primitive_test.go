package pq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinIntPQ(t *testing.T) {
	pq := NewMinIntPQ(4, 1, 3, 2)
	expected := []int{1, 2, 3, 4}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxIntPQ(t *testing.T) {
	pq := NewMaxIntPQ(4, 1, 3, 2)
	expected := []int{4, 3, 2, 1}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinInt64PQ(t *testing.T) {
	pq := NewMinInt64PQ(4, 1, 3, 2)
	expected := []int64{1, 2, 3, 4}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxInt64PQ(t *testing.T) {
	pq := NewMaxInt64PQ(4, 1, 3, 2)
	expected := []int64{4, 3, 2, 1}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinInt32PQ(t *testing.T) {
	pq := NewMinInt32PQ(4, 1, 3, 2)
	expected := []int32{1, 2, 3, 4}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxInt32PQ(t *testing.T) {
	pq := NewMaxInt32PQ(4, 1, 3, 2)
	expected := []int32{4, 3, 2, 1}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinInt16PQ(t *testing.T) {
	pq := NewMinInt16PQ(4, 1, 3, 2)
	expected := []int16{1, 2, 3, 4}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxInt16PQ(t *testing.T) {
	pq := NewMaxInt16PQ(4, 1, 3, 2)
	expected := []int16{4, 3, 2, 1}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinInt8PQ(t *testing.T) {
	pq := NewMinInt8PQ(4, 1, 3, 2)
	expected := []int8{1, 2, 3, 4}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxInt8PQ(t *testing.T) {
	pq := NewMaxInt8PQ(4, 1, 3, 2)
	expected := []int8{4, 3, 2, 1}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinFloat64PQ(t *testing.T) {
	pq := NewMinFloat64PQ(3.5, 2.1, 4.9, 1.2)
	expected := []float64{1.2, 2.1, 3.5, 4.9}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxFloat64PQ(t *testing.T) {
	pq := NewMaxFloat64PQ(3.5, 2.1, 4.9, 1.2)
	expected := []float64{4.9, 3.5, 2.1, 1.2}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinFloat32PQ(t *testing.T) {
	pq := NewMinFloat32PQ(3.5, 2.1, 4.9, 1.2)
	expected := []float32{1.2, 2.1, 3.5, 4.9}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxFloat32PQ(t *testing.T) {
	pq := NewMaxFloat32PQ(3.5, 2.1, 4.9, 1.2)
	expected := []float32{4.9, 3.5, 2.1, 1.2}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMinStringPQ(t *testing.T) {
	pq := NewMinStringPQ("delta", "alpha", "charlie", "bravo")
	expected := []string{"alpha", "bravo", "charlie", "delta"}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}

func TestNewMaxStringPQ(t *testing.T) {
	pq := NewMaxStringPQ("delta", "alpha", "charlie", "bravo")
	expected := []string{"delta", "charlie", "bravo", "alpha"}
	for _, v := range expected {
		assert.Equal(t, v, pq.Pop())
	}
	assert.True(t, pq.Empty())
}
