package main

import (
	"math/rand"

	"github.com/0-Mqix/YeetCode/solutions/heap"
)

//https://neetcode.io/

func main() {
	h := heap.New([]int{})

	for i := 0; i < 15; i++ {
		h.Push(rand.Intn(99-1) + 1)
		h.Print()
	}

	size := h.Size()

	for i := 0; i < size; i++ {
		h.Pop()
		h.Print()
	}
}
