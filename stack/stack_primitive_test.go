package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntStack(t *testing.T) {
	st := NewIntStack(1, 2, 3)
	assert.Equal(t, 3, st.Size())
	assert.False(t, st.IsEmpty())

	assert.Equal(t, 3, st.Top())
	_ = st.Pop()
	assert.Equal(t, 2, st.Top())
}

func TestInt8Stack(t *testing.T) {
	st := NewInt8Stack(10, 20)
	assert.Equal(t, int8(20), st.Top())

	_ = st.Pop()
	_ = st.Pop()
	assert.True(t, st.IsEmpty())
}

func TestInt16Stack(t *testing.T) {
	st := NewInt16Stack(100, 200)
	assert.Equal(t, int16(200), st.Top())
}

func TestInt32Stack(t *testing.T) {
	st := NewInt32Stack(300, 400)
	st.Push(500)
	assert.Equal(t, int32(500), st.Top())
}

func TestInt64Stack(t *testing.T) {
	st := NewInt64Stack(1000, 2000)
	assert.Equal(t, int64(2000), st.Top())
}

func TestFloat32Stack(t *testing.T) {
	st := NewFloat32Stack(1.1, 2.2)
	assert.Equal(t, float32(2.2), st.Top())
}

func TestFloat64Stack(t *testing.T) {
	st := NewFloat64Stack(3.3, 4.4)
	_ = st.Pop()
	assert.Equal(t, 3.3, st.Top())
}

func TestStringStack(t *testing.T) {
	st := NewStringStack("a", "b", "c")
	assert.Equal(t, "c", st.Top())
	st.Clear()
	assert.True(t, st.IsEmpty())
}

func TestRuneStack(t *testing.T) {
	st := NewRuneStack('x', 'y')
	assert.Equal(t, 'y', st.Top())
}

func TestByteStack(t *testing.T) {
	st := NewByteStack(0x1, 0x2)
	assert.Equal(t, byte(0x2), st.Top())
}

func TestStack_Clear(t *testing.T) {
	st := NewIntStack(10, 20, 30)
	assert.Equal(t, 3, st.Size())

	st.Clear()
	assert.Equal(t, 0, st.Size())
	assert.True(t, st.IsEmpty())
}

func TestStack_PopOnEmpty(t *testing.T) {
	st := NewIntStack()
	err := st.Pop()
	assert.Error(t, err)
}

func TestStack_TopOnEmpty(t *testing.T) {
	st := NewIntStack()
	top := st.Top()
	assert.Equal(t, 0, top) // zero value for int
}

func TestStack_Size(t *testing.T) {
	st := NewStringStack("hello", "world")
	assert.Equal(t, 2, st.Size())
	_ = st.Pop()
	assert.Equal(t, 1, st.Size())
}
