package solution1

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			if a+b == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
