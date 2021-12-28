package minheap

import "ds/constraints"

type node[T any, W constraints.Ordered] struct {
	element  T
	priority W
}

type nodes[T any, W constraints.Ordered] []node[T, W]

func (n nodes[T, W]) Less(i, j int) bool {
	return n[i].priority < n[j].priority
}

func (n nodes[T, W]) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
