package pq

// NewMinIntPQ creates a min-priority queue for integers.
func NewMinIntPQ(elements ...int) *PriorityQueue[int] {
	return NewPriorityQueue(func(a, b int) bool {
		return a < b
	}, elements...)
}

// NewMaxIntPQ creates a max-priority queue for integers.
func NewMaxIntPQ(elements ...int) *PriorityQueue[int] {
	return NewPriorityQueue(func(a, b int) bool {
		return a > b
	}, elements...)
}

// NewMinInt32PQ creates a min-priority queue for 32 bit integers.
func NewMinInt32PQ(elements ...int32) *PriorityQueue[int32] {
	return NewPriorityQueue(func(a, b int32) bool {
		return a < b
	}, elements...)
}

// NewMaxInt32PQ creates a max-priority queue for 32 bit integers.
func NewMaxInt32PQ(elements ...int32) *PriorityQueue[int32] {
	return NewPriorityQueue(func(a, b int32) bool {
		return a > b
	}, elements...)
}

// NewMinInt16PQ creates a min-priority queue for 16 bit integers.
func NewMinInt16PQ(elements ...int16) *PriorityQueue[int16] {
	return NewPriorityQueue(func(a, b int16) bool {
		return a < b
	}, elements...)
}

// NewMaxInt16PQ creates a max-priority queue for 16 bit integers.
func NewMaxInt16PQ(elements ...int16) *PriorityQueue[int16] {
	return NewPriorityQueue(func(a, b int16) bool {
		return a > b
	}, elements...)
}

// NewMinInt8PQ creates a min-priority queue for 8 bit integers.
func NewMinInt8PQ(elements ...int8) *PriorityQueue[int8] {
	return NewPriorityQueue(func(a, b int8) bool {
		return a < b
	}, elements...)
}

// NewMaxInt8PQ creates a max-priority queue for 8 bit integers.
func NewMaxInt8PQ(elements ...int8) *PriorityQueue[int8] {
	return NewPriorityQueue(func(a, b int8) bool {
		return a > b
	}, elements...)
}

// NewMinInt64PQ creates a min-priority queue for 64 bit integers.
func NewMinInt64PQ(elements ...int64) *PriorityQueue[int64] {
	return NewPriorityQueue(func(a, b int64) bool {
		return a < b
	}, elements...)
}

// NewMaxInt64PQ creates a max-priority queue for 64 bit integers.
func NewMaxInt64PQ(elements ...int64) *PriorityQueue[int64] {
	return NewPriorityQueue(func(a, b int64) bool {
		return a > b
	}, elements...)
}

// NewMinFloat64PQ creates a min-priority queue for float64 values.
func NewMinFloat64PQ(elements ...float64) *PriorityQueue[float64] {
	return NewPriorityQueue(func(a, b float64) bool {
		return a < b
	}, elements...)
}

// NewMaxFloat64PQ creates a max-priority queue for float64 values.
func NewMaxFloat64PQ(elements ...float64) *PriorityQueue[float64] {
	return NewPriorityQueue(func(a, b float64) bool {
		return a > b
	}, elements...)
}

// NewMinFloat32PQ creates a min-priority queue for float32 values.
func NewMinFloat32PQ(elements ...float32) *PriorityQueue[float32] {
	return NewPriorityQueue(func(a, b float32) bool {
		return a < b
	}, elements...)
}

// NewMaxFloat32PQ creates a max-priority queue for float32 values.
func NewMaxFloat32PQ(elements ...float32) *PriorityQueue[float32] {
	return NewPriorityQueue(func(a, b float32) bool {
		return a > b
	}, elements...)
}

// NewMinStringPQ creates a min-priority queue for strings (lexicographically smallest first).
func NewMinStringPQ(elements ...string) *PriorityQueue[string] {
	return NewPriorityQueue(func(a, b string) bool {
		return a < b
	}, elements...)
}

// NewMaxStringPQ creates a max-priority queue for strings (lexicographically largest first).
func NewMaxStringPQ(elements ...string) *PriorityQueue[string] {
	return NewPriorityQueue(func(a, b string) bool {
		return a > b
	}, elements...)
}
