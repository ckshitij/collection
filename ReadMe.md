# Go-Collection

**Go-Collection** is a **generic, type-safe, and efficient data structures library in Go**, inspired by popular collection frameworks like **C++ STL** and **Java Collections**.

It includes:

- **Priority Queue:** A heap-based priority queue supporting both min-heap and max-heap behaviors with generics.
- **Queue:** FIFO (First-In-First-Out) queue.
- **Stack:** LIFO (Last-In-First-Out) stack.
- **Set:** An unordered set of unique elements, similar to C++ std::unordered_set. Internally backed by a hash map for efficient insertions, deletions, and lookups, but does not maintain any ordering of elements.
- **List:** A **doubly linked list** with efficient front/back insertions and deletions.

> ✅ **Initial Release:** v0.1.0  
> **Go Version:** Go 1.18+ (for generics support)

---

## 📦 Installation

```bash
go get github.com/ckshitij/collection
```
