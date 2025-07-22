package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewList(t *testing.T) {
	l := NewList[int]()
	assert.NotNil(t, l)
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
	assert.Equal(t, 0, l.Len())
}

func TestPush(t *testing.T) {
	l := NewList[int]()
	l.PushFront(10)
	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.head, l.tail)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, 10, l.head.Element())

	l.PushFront(20)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 20, l.head.Element())
	assert.Equal(t, 10, l.tail.Element())
}

func TestAppend(t *testing.T) {
	l := NewList[int]()
	l.PushBack(10)
	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)
	assert.Equal(t, l.head, l.tail)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, 10, l.head.Element())

	l.PushBack(20)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 10, l.head.Element())
	assert.Equal(t, 20, l.tail.Element())
}

func TestPopFront(t *testing.T) {
	l := NewList[int]()
	l.PopFront() // No-op
	assert.Equal(t, (0), l.Len())

	l.PushFront(10)
	l.PopFront()
	assert.Equal(t, (0), l.Len())
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	l.PushFront(10)
	l.PushFront(20)
	l.PopFront()
	assert.Equal(t, (1), l.Len())
	assert.Equal(t, 10, l.head.Element())
}

func TestPopBack(t *testing.T) {
	l := NewList[int]()
	l.PopBack() // No-op
	assert.Equal(t, (0), l.Len())

	l.PushFront(10)
	l.PopBack()
	assert.Equal(t, (0), l.Len())
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	l.PushFront(10)
	l.PushFront(20)
	l.PopBack()
	assert.Equal(t, (1), l.Len())
	assert.Equal(t, 20, l.head.Element())
}

func TestFront(t *testing.T) {
	l := NewList[int]()
	assert.Nil(t, l.Front())

	l.PushFront(10)
	assert.NotNil(t, l.Front())
	assert.Equal(t, 10, l.Front().Element())
}

func TestBack(t *testing.T) {
	l := NewList[int]()
	assert.Nil(t, l.Back())

	l.PushBack(10)
	assert.NotNil(t, l.Back())
	assert.Equal(t, 10, l.Back().Element())
}

func TestLen(t *testing.T) {
	l := NewList[int]()
	assert.Equal(t, (0), l.Len())

	l.PushFront(10)
	assert.Equal(t, (1), l.Len())

	l.PushBack(20)
	assert.Equal(t, (2), l.Len())
}

func TestInsertAtPosition(t *testing.T) {
	l := NewList[int]()
	err := l.InsertAtPosition(10, -1)
	assert.Equal(t, ErrNegativePosition, err)

	err = l.InsertAtPosition(10, 0)
	require.NoError(t, err)
	assert.Equal(t, (1), l.Len())
	assert.Equal(t, 10, l.head.Element())

	err = l.InsertAtPosition(20, 1)
	require.NoError(t, err)
	assert.Equal(t, (2), l.Len())
	assert.Equal(t, 20, l.tail.Element())

	err = l.InsertAtPosition(30, 5)
	assert.Equal(t, ErrOutOfBound, err)

	err = l.InsertAtPosition(25, 1)
	require.NoError(t, err)
	assert.Equal(t, (3), l.Len())
	assert.Equal(t, 25, l.head.Next().Element())
}

func TestIterateForward(t *testing.T) {
	list := NewList[int]()
	results := []struct {
		index   int
		element int
	}{}

	// Case 1: Empty list
	list.IterateForward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Empty(t, 0, len(results)) // No elements to iterate

	// Case 2: Single element list
	list.PushBack(42)
	list.IterateForward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Len(t, results, 1)
	assert.Equal(t, 0, results[0].index)
	assert.Equal(t, 42, results[0].element)

	// Case 3: Multiple elements
	list.PushBack(100)
	list.PushBack(200)
	results = []struct {
		index   int
		element int
	}{} // Reset results
	list.IterateForward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Len(t, results, 3)
	assert.Equal(t, []struct {
		index   int
		element int
	}{
		{0, 42},
		{1, 100},
		{2, 200},
	}, results)
}

func TestIterateBackward(t *testing.T) {
	list := NewList[int]()
	results := []struct {
		index   int
		element int
	}{}

	// Case 1: Empty list
	list.IterateBackward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Empty(t, results) // No elements to iterate

	// Case 2: Single element list
	list.PushBack(42)
	list.IterateBackward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Len(t, results, 1)
	assert.Equal(t, (0), results[0].index)
	assert.Equal(t, 42, results[0].element)

	// Case 3: Multiple elements
	list.PushBack(100)
	list.PushBack(200)
	results = []struct {
		index   int
		element int
	}{} // Reset results
	list.IterateBackward(func(index int, element int) {
		results = append(results, struct {
			index   int
			element int
		}{index, element})
	})
	assert.Len(t, results, 3)
	assert.Equal(t, []struct {
		index   int
		element int
	}{
		{2, 200},
		{1, 100},
		{0, 42},
	}, results)
}
