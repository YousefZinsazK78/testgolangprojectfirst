package twosum

func TwoSum(nums []int, target int) []int {
	var result []int
	var firstIndex, lastIndex = 0, len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		_ = i
		if nums[firstIndex]+nums[lastIndex] == target {
			result = append(result, firstIndex, lastIndex)
			break
		}

	}
	return result
}
