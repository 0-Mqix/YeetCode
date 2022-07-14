package solutions

func ContainsDuplicate(nums []int) bool {
	hadNums := make(map[int]int)

	for _, v := range nums {
		hadNums[v]++
		if hadNums[v] > 1 {
			return true
		}
	}

	return false
}
