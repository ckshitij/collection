package queue

// NewIntQueue creates a new Queue for int values.
func NewIntQueue(elements ...int) *Queue[int] {
	q := NewQueue[int]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewInt8Queue creates a new Queue for int8 values.
func NewInt8Queue(elements ...int8) *Queue[int8] {
	q := NewQueue[int8]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewInt16Queue creates a new Queue for int16 values.
func NewInt16Queue(elements ...int16) *Queue[int16] {
	q := NewQueue[int16]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewInt32Queue creates a new Queue for int32 values.
func NewInt32Queue(elements ...int32) *Queue[int32] {
	q := NewQueue[int32]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewInt64Queue creates a new Queue for int64 values.
func NewInt64Queue(elements ...int64) *Queue[int64] {
	q := NewQueue[int64]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewFloat32Queue creates a new Queue for float32 values.
func NewFloat32Queue(elements ...float32) *Queue[float32] {
	q := NewQueue[float32]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewFloat64Queue creates a new Queue for float64 values.
func NewFloat64Queue(elements ...float64) *Queue[float64] {
	q := NewQueue[float64]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewStringQueue creates a new Queue for string values.
func NewStringQueue(elements ...string) *Queue[string] {
	q := NewQueue[string]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewRuneQueue creates a new Queue for rune values.
func NewRuneQueue(elements ...rune) *Queue[rune] {
	q := NewQueue[rune]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}

// NewByteQueue creates a new Queue for byte values.
func NewByteQueue(elements ...byte) *Queue[byte] {
	q := NewQueue[byte]()
	for _, e := range elements {
		q.Enqueue(e)
	}
	return q
}
