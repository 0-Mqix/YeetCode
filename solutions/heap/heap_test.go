package heap

import (
	"reflect"
	"testing"
)

func TestMinHeapInt(t *testing.T) {
	heap := New(make([]int, 0))
	values := []int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7}

	// Push
	for _, v := range values {
		heap.Push(v)
	}
	expected := []int{4, 6, 5, 8, 7, 10, 13, 12, 11, 9}
	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, got %v", expected, heap.data)
	}

	// Pop
	expected = []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for _, v := range expected {
		if val := heap.Pop(); val != v {
			t.Errorf("Expected %v, got %v", val, v)
		}
	}
}
