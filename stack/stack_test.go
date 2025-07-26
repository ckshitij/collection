package stack

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	// Test creating a new stack
	st := NewStack[int]()
	assert.NotNil(t, st)
	assert.True(t, st.IsEmpty())
}

func TestPush(t *testing.T) {
	st := NewStack[int]()

	// Test pushing elements
	st.Push(10)
	assert.False(t, st.IsEmpty())
	assert.Equal(t, 10, st.Top())

	st.Push(20)
	assert.False(t, st.IsEmpty())
	assert.Equal(t, 20, st.Top()) // Verify the stack's top element is updated
}

func TestIsEmpty(t *testing.T) {
	st := NewStack[string]()

	// Test an empty stack
	assert.True(t, st.IsEmpty())

	// Test a non-empty stack
	st.Push("Hello")
	assert.False(t, st.IsEmpty())

	// Test after removing all elements
	_ = st.Pop()
	assert.True(t, st.IsEmpty())
}

func TestPop(t *testing.T) {
	st := NewStack[int]()

	// Test popping from an empty stack
	err := st.Pop()
	require.EqualError(t, err, "invalid operation: empty stack")

	// Test popping from a non-empty stack
	st.Push(10)
	st.Push(20)
	err = st.Pop()
	require.NoError(t, err)
	assert.Equal(t, 10, st.Top()) // Verify the top element after pop

	// Test popping the last element
	err = st.Pop()
	require.NoError(t, err)
	assert.True(t, st.IsEmpty())
}

func TestTop(t *testing.T) {
	st := NewStack[int]()

	// Test top on an empty stack
	assert.Equal(t, 0, st.Top()) // Returns zero value for the type

	// Test top after pushing elements
	st.Push(1100000000000)
	assert.Equal(t, 1100000000000, st.Top())

	st.Push(1100000000001)
	assert.Equal(t, 1100000000001, st.Top()) // Top element without removing
}

func TestGenericStack(t *testing.T) {
	type customType struct {
		id   int
		name string
	}
	st := NewStack[customType]()

	// Test generic stack with custom type
	st.Push(customType{id: 1, name: "Alice"})
	st.Push(customType{id: 2, name: "Bob"})

	assert.Equal(t, customType{id: 2, name: "Bob"}, st.Top())

	err := st.Pop()
	require.NoError(t, err)
	assert.Equal(t, customType{id: 1, name: "Alice"}, st.Top())
}

func TestStackConcurrentAccess(t *testing.T) {
	stack := NewStack[int]()
	var wg sync.WaitGroup

	numGoroutines := 20
	opsPerGoroutine := 1000

	// Concurrent Push
	for i := range numGoroutines {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				stack.Push(base*opsPerGoroutine + j)
			}
		}(i)
	}

	// Concurrent Pop (won't block on empty, just ignore errors)
	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine/2; j++ {
				_ = stack.Pop()
			}
		}()
	}

	// Concurrent Top/Size/IsEmpty
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = stack.Top()
			_ = stack.Size()
			_ = stack.IsEmpty()
		}()
	}

	wg.Wait()

	// Final assertions
	require.GreaterOrEqual(t, stack.Size(), 0)

	if !stack.IsEmpty() {
		_ = stack.Top()
	}
}
