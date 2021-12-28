package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/samirettali/go-datastructures/minheap"
)

func main() {
	heap := minheap.New[int64, int64]()

	for i := 9; i >= 0; i-- {
		i := rand.Int63n(1000)
		heap.Push(i, i)
	}

	for heap.Size() > 0 {
		e, err := heap.Pop()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(e)
	}
}
