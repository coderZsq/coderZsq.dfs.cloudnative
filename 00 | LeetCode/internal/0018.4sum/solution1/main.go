package solution1

import (
	"fmt"
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
		if i > 0 && a == nums[i-1] { // 去重 a
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] { // 去重 b
				continue
			}
			for k := j + 1; k < len(nums)-1; k++ {
				c := nums[k]
				if k > j+1 && c == nums[k-1] { // 去重 c
					continue
				}
				for l := k + 1; l < len(nums); l++ {
					d := nums[l]
					if l > k+1 && d == nums[l-1] { // 去重 d
						continue
					}
					if a+b+c+d == target {
						result = append(result, []int{a, b, c, d})
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Print(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
