package solutions

//https://leetcode.com/problems/find-median-from-data-stream/submissions/

/*
Output
[null,null,-1.00000,null,-1.50000,null,-2.00000,null,-3.50000,null,-3.00000]
Expected
[null,null,-1.00000,null,-1.50000,null,-2.00000,null,-2.50000,null,-3.00000]
*/

type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{nums: make([]int, 0)}
}

func (m *MedianFinder) AddNum(num int) {
	m.nums = append(m.nums, num)
}

func (m *MedianFinder) FindMedian() float64 {
	l := len(m.nums)

	if l%2 != 0 {
		return float64(m.nums[l/2])
	}

	return float64(m.nums[l-2]+m.nums[l-1]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
