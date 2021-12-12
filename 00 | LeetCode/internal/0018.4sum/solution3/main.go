package solution3

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				c := nums[l]
				d := nums[r]
				sum := a + b + c + d
				if sum < target {
					l += 1
				} else if sum > target {
					r -= 1
				} else {
					result = append(result, []int{a, b, c, d})
					for l < r && c == nums[l+1] {
						l+=1
					}
					for l < r && d == nums[r-1] {
						r-=1
					}
					l+=1
					r-=1
				}
			}
		}
	}
	return result
}
