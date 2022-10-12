package solutions

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/find-median-from-data-stream/submissions/

// this was to slow :(, time to learn wat a heap map is :)
type MedianFinder struct {
	nums []int
	sort bool
}

func Constructor() MedianFinder {
	return MedianFinder{nums: make([]int, 0)}
}

func (m *MedianFinder) AddNum(num int) {
	m.nums = append(m.nums, num)
	m.sort = true
}

func (m *MedianFinder) FindMedian() float64 {
	if m.sort {
		sort.Ints(m.nums)
		m.sort = false
	}

	l := len(m.nums)
	p := l / 2

	if l%2 != 0 {
		fmt.Println("odd", m.nums)
		return float64(m.nums[p])
	}

	return float64(m.nums[p-1]+m.nums[p]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
