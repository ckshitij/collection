package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntQueue(t *testing.T) {
	q := NewIntQueue(1, 2, 3)
	assert.Equal(t, 3, q.Size())

	val := q.Front()
	assert.Equal(t, 1, val)

	val = q.Back()
	assert.Equal(t, 3, val)
}

func TestNewInt8Queue(t *testing.T) {
	q := NewInt8Queue(1, 2, 3)
	assert.Equal(t, 3, q.Size())

	err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, int8(2), q.Front())
}

func TestNewInt16Queue(t *testing.T) {
	q := NewInt16Queue(10, 20)
	assert.False(t, q.IsEmpty())
}

func TestNewInt32Queue(t *testing.T) {
	q := NewInt32Queue(100, 200)
	assert.Equal(t, int32(100), q.Front())
}

func TestNewInt64Queue(t *testing.T) {
	q := NewInt64Queue(1000, 2000)
	back := q.Back()
	assert.Equal(t, int64(2000), back)
}

func TestNewFloat32Queue(t *testing.T) {
	q := NewFloat32Queue(1.1, 2.2)
	assert.Equal(t, float32(1.1), q.Front())
}

func TestNewFloat64Queue(t *testing.T) {
	q := NewFloat64Queue(3.3, 4.4)
	err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 4.4, q.Front())
	assert.Equal(t, 4.4, q.Back())
}

func TestNewStringQueue(t *testing.T) {
	q := NewStringQueue("a", "b", "c")
	assert.Equal(t, "a", q.Front())
	assert.Equal(t, "c", q.Back())

	err := q.Dequeue()
	assert.NoError(t, err)
}

func TestNewRuneQueue(t *testing.T) {
	q := NewRuneQueue('x', 'y')
	assert.Equal(t, 'x', q.Front())
}

func TestNewByteQueue(t *testing.T) {
	q := NewByteQueue(0x1, 0x2)
	assert.Equal(t, byte(0x1), q.Front())
}

func TestQueue_EmptyDequeue(t *testing.T) {
	q := NewIntQueue()
	err := q.Dequeue()
	assert.Error(t, err)
}
