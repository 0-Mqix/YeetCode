package solutions

func TwoSum(nums []int, target int) []int {
	list := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		index, ok := list[target-nums[i]]

		if ok {
			return []int{i, index}
		}

		list[nums[i]] = i
	}

	return nil
}
