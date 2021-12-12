package solution3

import "sort"

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
		l, r := i+1, len(nums)-1
		for l < r {
			b := nums[l]
			c := nums[r]
			sum := a + b + c
			if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				result = append(result, []int{a, b, c})
				for l < r && b == nums[l+1] { // 去重 b
					l+=1
				}
				for l < r && c == nums[r-1] { // 去重 c
					r-=1
				}
				l+=1
				r-=1
			}
		}
	}
	return result
}