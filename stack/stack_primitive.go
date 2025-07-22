package stack

// NewIntStack creates a stack for int.
func NewIntStack(elements ...int) *Stack[int] {
	st := NewStack[int]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewInt8Stack creates a stack for int8.
func NewInt8Stack(elements ...int8) *Stack[int8] {
	st := NewStack[int8]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewInt16Stack creates a stack for int16.
func NewInt16Stack(elements ...int16) *Stack[int16] {
	st := NewStack[int16]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewInt32Stack creates a stack for int32.
func NewInt32Stack(elements ...int32) *Stack[int32] {
	st := NewStack[int32]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewInt64Stack creates a stack for int64.
func NewInt64Stack(elements ...int64) *Stack[int64] {
	st := NewStack[int64]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewFloat32Stack creates a stack for float32.
func NewFloat32Stack(elements ...float32) *Stack[float32] {
	st := NewStack[float32]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewFloat64Stack creates a stack for float64.
func NewFloat64Stack(elements ...float64) *Stack[float64] {
	st := NewStack[float64]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewStringStack creates a stack for string.
func NewStringStack(elements ...string) *Stack[string] {
	st := NewStack[string]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewRuneStack creates a stack for rune.
func NewRuneStack(elements ...rune) *Stack[rune] {
	st := NewStack[rune]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}

// NewByteStack creates a stack for byte.
func NewByteStack(elements ...byte) *Stack[byte] {
	st := NewStack[byte]()
	for _, e := range elements {
		st.Push(e)
	}
	return st
}
