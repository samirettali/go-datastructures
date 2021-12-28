package minheap

import (
	"fmt"

	"github.com/samirettali/go-datastructures/constraints"
)

var (
	ErrNoElements = fmt.Errorf("no elements")
)

// MinHeap struct implements a binary minimum heap using an array
type MinHeap[T any, W constraints.Ordered] struct {
	nodes nodes[T, W]
}

func New[T any, W constraints.Ordered]() *MinHeap[T, W] {
	return &MinHeap[T, W]{
		nodes: []node[T, W]{},
	}
}

func (h *MinHeap[T, W]) Push(element T, priority W) {
	n := node[T, W]{
		element:  element,
		priority: priority,
	}

	h.nodes = append(h.nodes, n)

	// Bubble up the inserted node
	i := len(h.nodes) - 1
	for i > 0 {
		p := parent(i)

		// If current element is greater or equal than it's parent, we're done
		if h.nodes[i].priority >= h.nodes[p].priority {
			return
		}

		// Swap element with it's parent
		h.nodes.Swap(i, p)

		i = p
	}
}

func (h *MinHeap[T, W]) Pop() (T, error) {
	if len(h.nodes) == 0 {
		return *new(T), ErrNoElements
	}

	// Get the root element of the heap
	node := h.nodes[0]

	// Replace the first element with the last one
	h.nodes[0] = h.nodes[len(h.nodes)-1]

	// Remove the last element
	h.nodes = h.nodes[:len(h.nodes)-1]

	// Bubble down the head
	i := 0
	for {
		min := i

		if l := left(i); l < len(h.nodes) && h.nodes.Less(l, min) {
			min = l
		}

		if r := right(i); r < len(h.nodes) && h.nodes.Less(r, min) {
			min = r
		}

		// If current element is smaller than both children, we're done
		if min == i {
			break
		}

		// Swap element with the smallest children
		h.nodes.Swap(i, min)

		i = min
	}

	return node.element, nil
}

func (h *MinHeap[T, W]) Peek() (T, error) {
	if len(h.nodes) == 0 {
		return *new(T), ErrNoElements
	}

	return h.nodes[0].element, nil
}

func (h *MinHeap[T, W]) Size() int {
	return len(h.nodes)
}

func parent(n int) int {
	return (n - 1) / 2
}

func left(n int) int {
	return n*2 + 1
}

func right(n int) int {
	return n*2 + 2
}
