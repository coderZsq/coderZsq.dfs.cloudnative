package solution1

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] { // 去重 a
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			b := nums[j]
			if j > i + 1 && b == nums[j-1] { // 去重 b
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				c := nums[k]
				if k > j + 1 && c == nums[k-1] { // 去重 c
					continue
				}
				if a+b+c == 0 {
					result = append(result, []int{a, b, c})
				}
			}
		}
	}
	return result
}