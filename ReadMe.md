# 📦 collection

A **Go (Golang) generic data structures library** featuring concurrency-safe and performant implementations of:

- ✅ Priority Queue (heap-based)
- ✅ Concurrency-safe Doubly Linked List
- ✅ Queue (FIFO)
- ✅ Stack (LIFO)
- ✅ Primitive-specific helper functions for performance and convenience

Built using Go **generics**, ensuring **type safety** without sacrificing performance.

---

## ✨ Features

### 1. **Priority Queue**
- Heap-based Min/Max Priority Queue.
- Supports **custom comparators**.
- Primitive-specific constructors:
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `float32`, `float64`
  - `string`

### 2. **List**
- Doubly Linked List implementation.
- **Concurrency-safe** using fine-grained locking.

### 3. **Queue**
- Generic FIFO queue built on top of the concurrency-safe list.
- Primitive-specific factory functions:
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `float32`, `float64`
  - `string`, `rune`, `byte`

### 4. **Stack**
- Generic LIFO stack built on top of the concurrency-safe list.
- Primitive-specific factory functions:
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `float32`, `float64`
  - `string`, `rune`, `byte`

### ✅ Common APIs
- `IsEmpty()`
- `Size()`
- `Clear()`
- Type-specific access methods (`Front`, `Back`, `Top`, `Pop`, `Push`)

---

## 📚 Usage

### Install

```bash
go get github.com/ckshitij/collection
```

### Example: Priority Queue

```go
import "github.com/ckshitij/collection/pq"

pq := pq.NewMinIntPQ(5, 3, 8)
pq.Push(2)
min := pq.Pop() // min == 2
```

### Example: Queue

```go
import "github.com/ckshitij/collection/queue"

q := queue.NewIntQueue(1, 2, 3)
q.Enqueue(4)
front := q.Front() // front == 1
_ = q.Dequeue()
```

### Example: Stack

```go
import "github.com/ckshitij/collection/stack"

st := stack.NewIntStack(10, 20, 30)
top := st.Top() // top == 30
_ = st.Pop()
```

### Example: Concurrency-Safe List

```go
import "github.com/ckshitij/collection/list"

lst := list.NewList[int]()
lst.PushFront(10)
lst.PushBack(20)
size := lst.Len()
```

---

## 🏗️ Contributing

We welcome contributions! To contribute:

1. **Fork** the repository.
2. Create a **feature/bugfix branch**.
3. Add your changes with appropriate **tests**.
4. Ensure all tests pass:

```bash
go test ./... -race -cover
```

5. Submit a **Pull Request** with a clear description.

### 📌 Contribution Guidelines
- Follow idiomatic Go practices.
- Ensure **>90% test coverage** for all new features.
- Maintain existing **naming conventions and structure**.

---

## 🔗 License

This project is licensed under the **MIT License**.

---

## 🔗 Repository

[GitHub Repository](https://github.com/ckshitij/collection)

